package response

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *JsonResponse) ToJson(c *gin.Context) {
	code := 200
	if res.Code > 50000 {
		code = 50000
	}
	c.JSON(code, res)
}

// ToJsonResponse 封装统一格式api返回的信息
func ToJsonResponse(code int, message string, data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data[0]
	}
	return &JsonResponse{
		Code:    code,
		Message: message,
		Data:    r,
	}
}
