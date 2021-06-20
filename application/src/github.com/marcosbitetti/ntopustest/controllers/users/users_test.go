package users

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var initialized = false
var router *gin.Engine

func prepareTest() *gin.Engine {
	if !initialized {
		initialized = true
		godotenv.Load("./../../../../../../../.env")
		router = gin.Default()
		Initialize(router)
	}
	return router
}

/*
 * Create a user
 */
func TestCreateUser(t *testing.T) {
	prepareTest()

	var jsonStr = []byte(`{
			"nome":"Testivaldo TESTANIO",
			"sexo":"M",
			"altura":1.78,
			"peso":58,
			"imc":18.31
		}\0`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/webapi/v1/user", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	req, _ = http.NewRequest("POST", "/webapi/v1/user", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	req, _ = http.NewRequest("POST", "/webapi/v1/user", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

/*func TestDuplicateUser(t *testing.T) {
	prepareTest()

	var jsonStr = []byte(`{
			"nome":"TestivaldoTESTANIO",
			"sexo":"M",
			"altura":1.78,
			"peso":58,
			"imc":18.31
		}\0`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/webapi/v1/user", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}*/

/*
 * Get user list
 */
/*func TestGetUsers(t *testing.T) {
	prepareTest()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/webapi/v1/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
*/
