package middleware

import (
	"APIGateway/constant"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenResp struct {
	UUID string `json:"uuid"`
}

// MiddleWare function
func MiddleWare(c *gin.Context) {
	params := struct {
		Token string `url:"token"`
	}{
		Token: c.Query("token"),
	}

	body, err := constant.RequestToService(constant.ServiceUser, http.MethodPost, "/", nil, params)
	if err != nil {
		constant.ResponseWithData(c, http.StatusForbidden, constant.PERMISSION_DENIED, nil)
		c.Abort()
		return
	}
	var tokenResp TokenResp
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		constant.ResponseWithData(c, http.StatusForbidden, constant.PERMISSION_DENIED, nil)
		c.Abort()
		return
	}
	c.Set("user_id", tokenResp.UUID)

	success := true
	if success {
		c.Next()
		return
	}

	c.Abort()
}
