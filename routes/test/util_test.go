package routes_test

import (
	"APIGateway/constant"
	"APIGateway/routes"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type test struct {
	name string
	path string
	args interface{}
	want constant.Response
}

func setUp() {
	constant.ReadConfig("../../.env")
	router = routes.InitRouter()
}

func tearDown() {
	
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setUp()
	exitCode := m.Run()
	tearDown()
	os.Exit(exitCode)
}
