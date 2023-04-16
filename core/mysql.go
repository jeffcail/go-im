package core

import (
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	RDB *redis.Client
)

func SetMysql(_db *gorm.DB) {
	Db = _db
}

func SetRedis(_rdb *redis.Client) {
	RDB = _rdb
}
