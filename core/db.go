package core

import (
	"github.com/streadway/amqp"
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	RDB *redis.Client
	MQ  *amqp.Connection
)

func SetMysql(_db *gorm.DB) {
	Db = _db
}

func SetRedis(_rdb *redis.Client) {
	RDB = _rdb
}

func SetRabbitMQ(_mq *amqp.Connection) {
	MQ = _mq
}
