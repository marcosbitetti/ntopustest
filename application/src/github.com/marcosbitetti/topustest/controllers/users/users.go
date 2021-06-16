package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marcosbitetti/topustest/models/user"
)

type FilteredItem struct {
	Nome   string
	Sexo   byte
	Peso   float32
	Altura float32
	IMC    float32
}

// GET /user
func GetUsers(c *gin.Context) {
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
	var name string = c.Param("name")
	var list []user.User = user.Find(name)
	if len(list) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": list[0]})
		return
	}
	//log.Printf("users ", len(user.Find("jose")))
	c.JSON(http.StatusOK, gin.H{"data": nil})
}

func Initialize(e *gin.Engine) {
	e.GET("/webapi/v1/user", GetUsers)
	e.GET("/webapi/v1/user/:name", GetUser)
}
