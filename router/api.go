package router

import (
	"net/http"

	auth2 "github.com/jeffcail/go-im/internel/http/handler/auth"

	"github.com/jeffcail/go-im/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterApiRouter 注册http路由服务
func RegisterApiRouter(router *gin.Engine) {

	router.Use(middlewares.StartCors())

	auth := new(auth2.AuthHandler)
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/send/register/email", auth.SendRegisterEmail)
		apiRouter.POST("/register", auth.Registered) // user register
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong...")
	})
}
