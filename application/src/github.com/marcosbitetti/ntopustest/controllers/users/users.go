package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marcosbitetti/ntopustest/models/user"
	"github.com/marcosbitetti/ntopustest/services/rabbitmq"
)

type FilteredItem struct {
	Nome   string
	Sexo   byte
	Peso   float32
	Altura float32
	IMC    float32
}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

// GET /user
func GetUsers(c *gin.Context) {
	cors(c)
	// filter
	tmp := user.FindAll()
	var list []FilteredItem = make([]FilteredItem, len(tmp))
	for i, u := range tmp {
		list[i] = FilteredItem{u.Nome, u.Sexo, u.Altura, u.Peso, u.IMC}
	}
	//log.Printf("users ", len(user.Find("jose")))
	c.JSON(http.StatusOK, gin.H{"data": list})
}

// GET /user/:name
func GetUser(c *gin.Context) {
	cors(c)
	var name string = c.Param("name")
	var list []user.User = user.Find(name)
	if len(list) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": list[0]})
		return
	}
	//log.Printf("users ", len(user.Find("jose")))
	c.JSON(http.StatusOK, gin.H{"data": nil})
}

type UserReqeust struct {
	Nome   string  `json:"nome" binding:"required"`
	Sexo   string  `json:"sexo" binding:"required"`
	Altura float32 `json:"altura" binding:"required"`
	Peso   float32 `json:"peso" binding:"required"`
	IMC    float32 `json:"imc" binding:"required"`
}

// POST /user
func CreateUser(c *gin.Context) {
	cors(c)
	var u UserReqeust

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong data"})
		return
	}

	n, ok := user.New(u.Nome, user.MapSexo[u.Sexo], u.Altura, u.Peso, u.IMC)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user-exists"})
		return
	}

	n.Commit()
	rabbitmq.Broadcaster <- "CADASTRO novo usuário \"" + n.Nome + "\""
	c.JSON(http.StatusCreated, gin.H{"data": n})
}

// Put /user
func UpdateUser(c *gin.Context) {
	cors(c)
	var u UserReqeust

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong data"})
		return
	}

	var name string = c.Param("name")
	var list []user.User = user.Find(name)
	if len(list) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not-found"})
		return
	}

	up := list[0]
	message, ok := up.Update(u.Nome, user.MapSexo[u.Sexo], u.Altura, u.Peso, u.IMC)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	rabbitmq.Broadcaster <- "ATUALIZAÇÃO nos registros do usuário \"" + up.Nome + "\""
	c.JSON(http.StatusCreated, gin.H{"data": up})
}

// DELETE /user/:name
func DeleteUser(c *gin.Context) {
	cors(c)
	var name string = c.Param("name")
	var list []user.User = user.Find(name)

	if len(list) > 0 {
		u := list[0]
		if !u.Delete() {
			c.JSON(http.StatusBadRequest, gin.H{"message": "internal-error"})
			return
		}

		rabbitmq.Broadcaster <- "EXCLUSÃO do usuário \"" + u.Nome + "\""
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "internal-error"})
}

func Initialize(e *gin.Engine) {
	e.GET("/webapi/v1/user", GetUsers)
	e.GET("/webapi/v1/user/:name", GetUser)
	e.POST("/webapi/v1/user", CreateUser)
	e.PUT("/webapi/v1/user/:name", UpdateUser)
	e.DELETE("/webapi/v1/user/:name", DeleteUser)
}
