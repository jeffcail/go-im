package boot

import (
	"github.com/jeffcail/go-im/config"
	"github.com/jeffcail/go-im/core"
	_init "github.com/jeffcail/go-im/init"
)

// Run 项目启动初始化工作
func Run() {
	// 解析配置
	config.InitParse()
	// 初始化数据库连接
	db := _init.InitMysql()
	core.SetMysql(db)
	// 初始化redis连接

	rdb := _init.InitRedis()
	core.SetRedis(rdb)
}