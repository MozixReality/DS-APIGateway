package constant

import (
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
	ERROR_HANDLE_MEDIA
	ERROR_TOKEN_CLAIMS_PARSING_FAILED
)

var msgFlags = map[int]string{
	SUCCESS:                           "Ok",
	INVALID_PARAMS:                    "Invalid params error",
	ERROR:                             "Fail",
	ERROR_DATABASE:                    "Error retrieving from database",
	ERROR_HANDLE_MEDIA:                "Error handling media upload to CDN",
	ERROR_TOKEN_CLAIMS_PARSING_FAILED: "Failed to parse token claims",
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
