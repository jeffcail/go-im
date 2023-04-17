package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/go-im/middlewares"
	"github.com/jeffcail/go-im/router"
)

// ApiServer 启动http api 服务
func ApiServer() {
	r := gin.Default()

	r.Use(middlewares.Recover)

	router.RegisterApiRouter(r)

	r.Run(":8080")
}
