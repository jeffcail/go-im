package _init

import (
	"github.com/jeffcail/go-im/global"
	"github.com/jeffcail/go-im/pkg/db"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	db, err := db.InitMysqlConnect()
	global.CheckErr(err)
	return db
}
