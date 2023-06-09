package config

import (
	"log"
	"os"

	"github.com/jeffcail/go-im/global"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	AppName       string `json:"app_name"`
	AppUrl        string `json:"app_url"`
	AppGoroutines int    `json:"app_goroutines"`
	HTTPBind      string `json:"http_bind"`
	Mysql         struct {
		DbDsn       string `json:"db_dsn"`
		MaxOpenConn int    `json:"max_open_conn"`
		MaxIdleConn int    `json:"max_idle_conn"`
	} `json:"mysql"`
	Redis struct {
		RedisAddr string `json:"redis_addr"`
		Password  string `json:"password"`
		RedisDb   int    `json:"redis_db"`
	} `json:"redis"`
	Email struct {
		MailFrom       string `json:"mail_from"`
		MailPassword   string `json:"mail_password"`
		MailServer     string `json:"mail_server"`
		MailServerPort string `json:"mail_server_port"`
	} `json:"email"`
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
		Expire int64  `json:"expire"`
		SECRET string `json:"secret"`
	} `json:"jwt"`
	SmMs struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Token    string `json:"token"`
	} `json:"sm_ms"`
}

var Config *GlobalConfig

// InitParse 解析项目配置文件
func InitParse() {
	viper.SetConfigFile("./go-im.yaml")
	err := viper.ReadInConfig()
	global.CheckErr(err)
	err = viper.Unmarshal(&Config)
	Config.Email.MailFrom = os.Getenv("Mail163From")
	Config.Email.MailPassword = os.Getenv("Mail163Pass")
	if Config.Email.MailFrom == "" || Config.Email.MailPassword == "" {
		log.Fatal("配置信息不全,请完善邮箱相关配置信息")
	}
	Config.SmMs.Name = os.Getenv("SMMSName")
	Config.SmMs.Password = os.Getenv("SMMSPassword")
	Config.SmMs.Token = os.Getenv("SMMSToken")
	if Config.SmMs.Name == "" || Config.SmMs.Password == "" || Config.SmMs.Token == "" {
		log.Fatal("配置信息不全,请完善sm.ms相关配置信息")
	}
	global.CheckErr(err)
}
