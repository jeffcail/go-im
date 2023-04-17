package logr

import (
	"github.com/jeffcail/go-im/config"
	"github.com/jeffcail/wildrocket/wlog"
	"go.uber.org/zap"
)

// InitLogger 初始化日志
func InitLogger() *zap.Logger {
	wlog.InitLogger(config.Config.Logger.Path)
	return wlog.Wlogger
}
