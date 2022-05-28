package v1

import (
	"APIGateway/constant"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type LineAccount struct {
	Nickname      string `json:"nick_name" binding:"required" example:"mozix"`
	AccessToken   string `json:"access_token" binding:"required" example:"+ArlB6Bg1AJjq8rT6hzn0jVd5JLOpWOy3rYoLUg/GuyfgzyKMImccgXuyAvI7Mhyxr9/NH7opIcxP/0D1s4/LbuvNDSFQAlPPrQOot8wPWlnI1sO51R+ugG6Onsf6fo/faarHspzWABrV3QckCZeWwdB04t89/1O/w1cDnyilFU="`
	ChannelSecret string `json:"channel_secret" binding:"required" example:"05425dffa740b7c5d81f89fc5993e074"`
}

type LineMessage struct {
	Nickname string `json:"nick_name" example:"mozix"`
	Message  string `json:"message" example:"Hello there~"`
}

// @Summary Get Line Info
// @produce application/json
// @param token query string true "Authorization"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /line_info [get]
func GetLineInfo(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		constant.ResponseWithData(c, http.StatusForbidden, constant.PERMISSION_DENIED, nil)
		return
	}
	params := struct {
		UserID string `url:"user_id"`
	}{
		UserID: userID.(string),
	}

	body, err := constant.RequestToService(constant.ServiceLine, http.MethodGet, "/line_info", c.Request.Body, params)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error())
		return
	}

	constant.ResponseWithBody(c, http.StatusOK, body)
}

// @Summary Update Line Info
// @produce application/json
// @param token query string true "Authorization"
// @Param LineAccount body LineAccount true "Update Line Info"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /line_info [put]
func UpdateLineInfo(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		constant.ResponseWithData(c, http.StatusForbidden, constant.PERMISSION_DENIED, nil)
		return
	}
	var lineAccount LineAccount
	if err := c.ShouldBindJSON(&lineAccount); err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}

	params := struct {
		UserID        string `url:"user_id"`
		AccessToken   string `url:"access_token"`
		ChannelSecret string `url:"channel_secret"`
		NickName      string `url:"nick_name"`
	}{
		UserID:        userID.(string),
		AccessToken:   lineAccount.AccessToken,
		ChannelSecret: lineAccount.ChannelSecret,
		NickName:      lineAccount.Nickname,
	}

	body, err := constant.RequestToService(constant.ServiceLine, http.MethodPut, "/line_info", nil, params)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error())
		return
	}

	constant.ResponseWithBody(c, http.StatusOK, body)
}

// @Summary Send Line Message
// @produce application/json
// @param token query string true "Authorization"
// @Param LineMessage body LineMessage true "Send Line Message"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /send_message [post]
func SendLineMessage(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		constant.ResponseWithData(c, http.StatusForbidden, constant.PERMISSION_DENIED, nil)
		return
	}

	var lineMessage LineMessage
	if err := c.ShouldBindJSON(&lineMessage); err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.INVALID_PARAMS, err.Error())
		return
	}

	writer := kafka.Writer{
		Addr:     kafka.TCP(viper.GetString("KAFKA_URL")),
		Topic:    viper.GetString("TOPIC"),
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	msg := kafka.Message{
		Value: []byte(fmt.Sprintf("%s==========%s==========%s", userID, lineMessage.Nickname, lineMessage.Message)),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.ERROR, err.Error())
		return
	}

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, nil)
}
