package pool

import (
	"github.com/jeffcail/go-im/config"
	"github.com/panjf2000/ants/v2"
)

var Gpoll *ants.Pool

// BootGoroutinePool 启动协程池管理
func BootGoroutinePool() {
	Gpoll, _ = ants.NewPool(config.Config.AppGoroutines)
}
