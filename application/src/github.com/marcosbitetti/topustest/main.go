package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/marcosbitetti/topustest/controllers/users"
)

var indexContent string

func IndexPage(c *gin.Context) {
	if indexContent == "" {
		// prepare default response
		_indexContent, err := ioutil.ReadFile("./views/index.html")
		if err != nil {
			panic(err)
		}
		indexContent = string(_indexContent)
	}
	c.String(http.StatusOK, indexContent)
}

func main() {
	// check if environment is loaded. if not, assume that is on developer
	if os.Getenv("DB_NAME") == "" {
		err := godotenv.Load("./../../../../../.env")
		if err != nil {
			panic(err)
		}
	}

	/*database.DB()
	// database.UserCollection()
	u := user.New("ailoze", user.MALE, 1.6, 68.0, 20.0)
	//bug, _ :=
	database.Insert(u, database.UserCollection())
	list := database.FindAll(&user.User{}, database.UserCollection())
	log.Println("lista")
	log.Println(len(list))
	var list2 []user.User = make([]user.User, len(list))
	log.Println(len(list2))
	for i, _u := range list {
		u, ok := _u.(*user.User)
		if ok {
			log.Println(u.Nome)
			log.Println(u.Id)
			list2[i] = *u
		} else {
			log.Println("err: ", u.Nome)
		}
	}*/

	web := gin.Default()

	web.GET("/", IndexPage)

	users.Initialize(web)

	log.Printf("Rodando servidor web na porta %s", os.Getenv("APPLICATION_PORT"))

	web.Run(":" + os.Getenv("APPLICATION_PORT"))
}
