package constant

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS        = http.StatusOK
	INVALID_PARAMS = http.StatusBadRequest
	ERROR          = http.StatusInternalServerError
)

const (
	ERROR_DATABASE = iota + 10000
	PERMISSION_DENIED
)

var msgFlags = map[int]string{
	SUCCESS:           "Ok",
	INVALID_PARAMS:    "Invalid params error",
	ERROR:             "Fail",
	ERROR_DATABASE:    "Error retrieving from database",
	PERMISSION_DENIED: "Permission denied",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(c *gin.Context, httpCode, respCode int, data interface{}) {
	msg, ok := msgFlags[respCode]
	if !ok {
		msg = msgFlags[ERROR]
	}
	response := Response{
		Code:    respCode,
		Message: msg,
		Data:    data,
	}
	c.JSON(httpCode, response)
}

func ResponseWithBody(c *gin.Context, httpCode int, body []byte) {
	var response Response
	var data interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		response.Code = ERROR
	}
	response.Code = http.StatusOK
	response.Message = "Ok"
	response.Data = data
	c.JSON(httpCode, response)
}
