package _init

import (
	"github.com/jeffcail/go-im/global"
	"github.com/jeffcail/go-im/pkg/mq"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() *amqp.Connection {
	mq, err := mq.InitMq()
	global.CheckErr(err)
	return mq
}
