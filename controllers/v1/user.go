package v1

import (
	"APIGateway/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginResp struct {
	Token 	string	`json:"token" binding:"required"`
	Expire 	int		`json:"expire" binding:"required"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required" example:"mozixreality"`
	Password string `json:"password" binding:"required" example:"Abcd1234"`
}

// @Summary Sign up
// @produce application/json
// @Param LoginReq body LoginReq true "Sign up"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /sign_up [post]
func SignUp(c *gin.Context) {
	loginResp := LoginResp{
		Token: "quWIU3287fnU",
		Expire: 7200,
	}
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, loginResp)
}

// @Summary Sign in
// @produce application/json
// @Param LoginReq body LoginReq true "Sign in"
// @Success 200 {string} constant.Response
// @Failure 500 {object} constant.Response
// @Router /sign_in [post]
func SignIn(c *gin.Context) {
	loginResp := LoginResp{
		Token: "quWIU3287fnU",
		Expire: 7200,
	}
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, loginResp)
}