package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

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

	web := gin.Default()

	web.GET("/", IndexPage)

	users.Initialize(web)

	log.Printf("Rodando servidor web na porta %s", os.Getenv("APPLICATION_PORT"))

	web.Run(":" + os.Getenv("APPLICATION_PORT"))
}
