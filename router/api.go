package router

import (
	"net/http"

	"github.com/jeffcail/go-im/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterApiRouter 注册http路由服务
func RegisterApiRouter(router *gin.Engine) {

	router.Use(middlewares.StartCors())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong...")
	})
}
