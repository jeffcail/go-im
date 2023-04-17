package _init

import (
	logr "github.com/jeffcail/go-im/pkg/log"
	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	return logr.InitLogger()
}
