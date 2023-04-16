package config

import (
	"github.com/jeffcail/go-im/global"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	AppName       string `json:"app_name"`
	AppUrl        string `json:"app_url"`
	AppGoroutines string `json:"app_goroutines"`
	HTTPBind      string `json:"http_bind"`
	Mysql         struct {
		DbDsn string `json:"db_dsn"`
	} `json:"mysql"`
	Redis struct {
		RedisAddr string `json:"redis_addr"`
		Password  string `json:"password"`
		RedisDb   string `json:"redis_db"`
	} `json:"redis"`
	RabbitMQ struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"rabbit_mq"`
	Logger struct {
		Path      string `json:"path"`
		MaxSize   string `json:"max_size"`
		MaxAge    string `json:"max_age"`
		Compress  string `json:"compress"`
		LocalTime string `json:"local_time"`
	} `json:"logger"`
	Jwt struct {
		Expire string `json:"expire"`
		SECRET string `json:"secret"`
	} `json:"jwt"`
}

var Config *GlobalConfig

// InitParse 解析项目配置文件
func InitParse() {
	viper.SetConfigFile("./go-im.yaml")
	err := viper.ReadInConfig()
	global.CheckErr(err)
	err = viper.Unmarshal(&Config)
	global.CheckErr(err)
}
