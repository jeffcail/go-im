package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/jeffcail/go-im/core"

	"github.com/gin-gonic/gin"
)

// 全局异常处理
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			er := errToStr(r)
			core.ImLogger.Warn(er)
			c.JSONP(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器异常",
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

func errToStr(err interface{}) string {
	switch v := err.(type) {
	case error:
		return v.Error()
	default:
		return err.(string)
	}
}
