package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//订制配置文件解析载体
type config struct {
	Web        *webset
	Database   *database
	Redis      *redis
	Alisms     *alisms
	Jwt        *jwtInfo
	CaptchaDir *captchaDirInfo
}

type captchaDirInfo struct {
	Path string
}
type jwtInfo struct {
	Key string
}

type alisms struct {
	Accessid  string
	Accesskey string
	Signname  string
	Template  string
}

type redis struct {
	//Driver   string
	Poolsize int
	Host     string
	Port     uint
	Password string
}

//订制Database块
type database struct {
	//Driver   string
	Poolsize int
	Host     string
	Port     uint
	Username string
	Dbname   string
	Password string
}

type webset struct {
	Model string
	Port  string
}

var Config *config = new(config)

func init() {
	//读取配置文件
	_, err := toml.DecodeFile("config.toml", Config)
	if err != nil {
		log.Panic(err.Error())
	}
}
