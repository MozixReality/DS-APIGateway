package v1_test

import (
	"APIGateway/constant"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

type test struct {
	name string
	args interface{}
	want interface{}
}

func setUp() {
	constant.ReadConfig("../../../.env")
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
