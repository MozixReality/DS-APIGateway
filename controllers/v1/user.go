package v1

import (
	"APIGateway/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" url:"username" binding:"required" example:"mozixreality"`
	Password string `json:"password" url:"password" binding:"required" example:"Abcd1234"`
}

// @Summary Sign up
// @produce application/json
// @Param LoginReq body LoginReq true "Sign up"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /sign_up [post]
func SignUp(c *gin.Context) {
	var loginReq LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.INVALID_PARAMS, err.Error())
		return
	}
	body, err := constant.RequestToService(constant.ServiceUser, http.MethodPost, "/sign_up", nil, loginReq)
	if err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.ERROR, err.Error())
		return
	}
	constant.ResponseWithBody(c, http.StatusOK, body)
}

// @Summary Sign in
// @produce application/json
// @Param LoginReq body LoginReq true "Sign in"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /sign_in [post]
func SignIn(c *gin.Context) {
	var loginReq LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.INVALID_PARAMS, err.Error())
		return
	}
	body, err := constant.RequestToService(constant.ServiceUser, http.MethodPost, "/sign_in", nil, loginReq)
	if err != nil {
		constant.ResponseWithData(c, http.StatusBadRequest, constant.ERROR, err.Error())
		return
	}
	constant.ResponseWithBody(c, http.StatusOK, body)
}
