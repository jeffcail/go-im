package db

import (
	"github.com/jeffcail/go-im/config"
	"github.com/jeffcail/wildrocket/wdb"
	"gorm.io/gorm"
)

// InitMysqlConnect 项目启动 初始化mysql连接
func InitMysqlConnect() (db *gorm.DB, err error) {
	return wdb.InitGormMysql(config.Config.Mysql.DbDsn,
		[]int{config.Config.Mysql.MaxOpenConn, config.Config.Mysql.MaxIdleConn})
}
