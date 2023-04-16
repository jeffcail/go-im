package _init

import (
	"github.com/jeffcail/go-im/global"
	"github.com/jeffcail/go-im/pkg/redis"
	redis2 "gopkg.in/redis.v5"
)

func InitRedis() *redis2.Client {
	client, err := redis.InitRedisConnect()
	global.CheckErr(err)
	return client
}
