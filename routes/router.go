package routes

import (
	"APIGateway/controllers"
	v1 "APIGateway/controllers/v1"
	_ "APIGateway/docs"
	"APIGateway/middleware"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AddAllowMethods(http.MethodOptions)
	config.AddAllowHeaders("Authorization", "Upgrade", "Connection", "Accept", "Accept-Encoding", "Accept-Language", "Host", "Cookie", "Referer", "User-Agent")
	config.AddExposeHeaders("Token")
	config.AllowCredentials = true
	config.AllowWildcard = true
	config.AllowOriginFunc = func(origin string) bool {
		allowOriginDomains := []string{"localhost", "127.0.0.1", "aralego.com", "showmore.cc"}
		for _, d := range allowOriginDomains {
			if strings.Contains(origin, d) {
				return true
			}
		}
		return false
	}
	if err := config.Validate(); err != nil {
		panic(err)
	}
	router.Use(cors.New(config))

	if mode := gin.Mode(); mode == gin.DebugMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.GET("/heartBeat", controllers.HeartBeat)

	router.POST("/sign_up", v1.SignUp)
	router.POST("/sign_in", v1.SignIn)

	router.Use(middleware.MiddleWare)

	router.GET("/line_info", v1.GetLineInfo)
	router.PUT("/line_info", v1.UpdateLineInfo)
	router.POST("/send_message", v1.SendLineMessage)


	return router
}
