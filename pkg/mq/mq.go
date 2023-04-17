package mq

import (
	"github.com/jeffcail/go-im/config"
	"github.com/jeffcail/wildrocket/wmq"
	"github.com/streadway/amqp"
)

// InitMq 初始化mq连接
func InitMq() (mq *amqp.Connection, err error) {
	return wmq.InitRabbitMQClient(config.Config.RabbitMQ.User, config.Config.RabbitMQ.Password,
		config.Config.RabbitMQ.Host, config.Config.RabbitMQ.Port)
}
