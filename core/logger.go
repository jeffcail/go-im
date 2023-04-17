package core

import "go.uber.org/zap"

var ImLogger *zap.Logger

func SetLogger(_logger *zap.Logger) {
	ImLogger = _logger
}
