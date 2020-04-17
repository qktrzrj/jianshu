package setting

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Storage interface {
	GetHost() string
	GetPort() int
	GetUsername() string
	GetPassword() string
	GetDBName() string
}

type storageConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func (s storageConfig) GetHost() string {
	return s.Host
}

func (s storageConfig) GetPort() int {
	return s.Port
}

func (s storageConfig) GetUsername() string {
	return s.Username
}

func (s storageConfig) GetPassword() string {
	return s.Password
}

func (s storageConfig) GetDBName() string {
	return s.DBName
}

var (
	jwtSecret string
	httpPort  int
	storage   storageConfig
	once      sync.Once
)

func Init() {
	once.Do(func() {
		viper.AddConfigPath("config")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("读取配置文件出错：%s", err)
		}

		//设置默认值
		viper.SetDefault("app.jwt_secret", "secure")
		viper.SetDefault("http.port", 8008)

		// 读取配置内容
		jwtSecret = viper.GetString("app.jwt_secret")
		httpPort = viper.GetInt("http.port")
		err = viper.UnmarshalKey("storage", &storage, func(config *mapstructure.DecoderConfig) {
			config.TagName = "json"
		})
		if err != nil {
			log.Fatalf("读取数据库配置信息出错：%s", err)
		}

		//// 监控JwtSecret
		//go func() {
		//	for {
		//		viper.WatchConfig()
		//		jwtSecret = viper.GetString("app.jwt_secret")
		//	}
		//}()
	})
}

func GetJwtSecret() string {
	return jwtSecret
}

func GetHttpPort() int {
	return httpPort
}

func GetStorage() Storage {
	return storage
}
