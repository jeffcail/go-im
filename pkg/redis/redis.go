package redis

import (
	"github.com/jeffcail/go-im/config"
	"github.com/jeffcail/wildrocket/wdb"
	"gopkg.in/redis.v5"
)

// InitRedisConnect redis连接
func InitRedisConnect() (client *redis.Client, err error) {
	return wdb.InitRedisClient(config.Config.Redis.RedisAddr, config.Config.Redis.Password, config.Config.Redis.RedisDb)
}
