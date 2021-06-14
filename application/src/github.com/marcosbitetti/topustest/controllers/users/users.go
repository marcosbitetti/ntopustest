package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marcosbitetti/topustest/models/user"
)

// GET /user
func Users(c *gin.Context) {
	u := user.New("jose", '1', 1.80, 68.0, 20.98765)
	log.Printf("users ", len(user.Find("jose")))
	c.JSON(http.StatusOK, gin.H{"data": u.Nome()})
}

func Initialize(e *gin.Engine) {
	e.GET("/webapi/v1/user", Users)
}
