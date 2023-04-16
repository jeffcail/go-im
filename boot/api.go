package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/go-im/router"
)

// ApiServer 启动http api 服务
func ApiServer() {
	imApi := gin.Default()

	router.RegisterApiRouter(imApi)

	imApi.Run(":8080")
}
