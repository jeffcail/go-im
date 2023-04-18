package im

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"

	"github.com/jeffcail/go-im/config"

	"github.com/jeffcail/go-im/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/jeffcail/go-im/core"
)

type (
	SmApiHandler      struct{}
	TokenResponseData struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
		Success   bool   `json:"success"`
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestId string `json:"request_id"`
	}
)

const (
	CacheSmToken = "sm-token"

	SmGetApiTokenUrl = "https://sm.ms/api/v2/token" // 获取sm.ms token
)

// GetApiToken 获取sm.ms token
// https://doc.sm.ms/#api-Image-Upload
// User - Get API-Token
func (sm *SmApiHandler) GetApiToken(c *gin.Context) {
	val := core.RDB.Get(CacheSmToken).Val()
	if val != "" {
		r := &TokenResponseData{}
		r.Data.Token = val
		r.Code = "success"
		r.Success = true

		response.ToJsonResponse(response.Success, "获取api token 成功", r).ToJson(c)
		return
	}
	values := url.Values{"username": {config.Config.SmMs.Name}, "password": {config.Config.SmMs.Password}}
	resp, err := http.PostForm(SmGetApiTokenUrl, values)
	if err != nil {
		core.ImLogger.Error("获取sm.ms token failed", zap.Error(err))
		response.ToJsonResponse(response.GetSmMsApiTokenFailed, "获取api token 失败").ToJson(c)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		core.ImLogger.Error("获取sm.ms token failed", zap.Error(err))
		response.ToJsonResponse(response.GetSmMsApiTokenFailed, "获取api token 失败").ToJson(c)
		return
	}
	res := &TokenResponseData{}
	if err = json.Unmarshal(body, res); err != nil {
		core.ImLogger.Error("获取sm.ms token unmarshal json failed", zap.Error(err))
	}
	if !res.Success {
		response.ToJsonResponse(http.StatusInternalServerError, res.Message).ToJson(c)
		return
	}
	if err = core.RDB.Set(CacheSmToken, res.Data.Token, time.Hour*1).Err(); err != nil {
		core.ImLogger.Error("获取sm.ms token success write in cache failed", zap.Error(err))
	}

	response.ToJsonResponse(response.Success, "获取sm api token成功", res).ToJson(c)
	return
}
