package v1

import (
	"APIGateway/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LineAccount struct {
	Nickname 		string 	`json:"nick_name" binding:"required" example:"mozix"`
	ChannelID 		string 	`json:"channel_id" binding:"required" example:"1656452143"`
	ChannelSecret 	string 	`json:"channel_secret" binding:"required" example:"05425dffa740b7c5d81f89fc5993e074"`
}

type LineMessage struct {
	Nickname 	string 	`json:"nick_name" example:"mozix"`
	Message 	string 	`json:"message" example:"Hello there~"`
}

// @Summary Get Line Info
// @produce application/json
// @param token query string true "Authorization"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /line_info [get]
func GetLineInfo(c *gin.Context) {
	LineAccounts := []LineAccount{}
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, LineAccounts)
}

// @Summary Update Line Info
// @produce application/json
// @param token query string true "Authorization"
// @Param LineAccount body LineAccount true "Update Line Info"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /line_info [put]
func UpdateLineInfo(c *gin.Context) {
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, nil)
}

// @Summary Send Line Message
// @produce application/json
// @param token query string true "Authorization"
// @Param LineMessage body LineMessage true "Send Line Message"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /send_message [post]
func SendLineMessage(c *gin.Context) {
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, nil)
}