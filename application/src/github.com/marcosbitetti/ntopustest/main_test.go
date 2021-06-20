package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/*
 * IndexPage return Status 200 if server is ok
 * content is ignored at first time because it is not
 * have a specification on project
 */
func TestIndexPage(t *testing.T) {
	router := gin.Default()
	router.GET("/", IndexPage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
