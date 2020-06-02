package setting

import (
	"github.com/fsnotify/fsnotify"
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

type Redis interface {
	GetHost() string
	GetPort() int
	GetDB() int
	GetPassword() string
}

type storageConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type redisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func (s storageConfig) GetHost() string     { return s.Host }
func (s storageConfig) GetPort() int        { return s.Port }
func (s storageConfig) GetUsername() string { return s.Username }
func (s storageConfig) GetPassword() string { return s.Password }
func (s storageConfig) GetDBName() string   { return s.DBName }

func (r redisConfig) GetHost() string     { return r.Host }
func (r redisConfig) GetPort() int        { return r.Port }
func (r redisConfig) GetDB() int          { return r.DB }
func (r redisConfig) GetPassword() string { return r.Password }

var (
	jwtSecret string
	httpPort  int
	esUrl     string
	storage   storageConfig
	redis     redisConfig
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
		esUrl = viper.GetString("es.url")
		err = viper.UnmarshalKey("storage", &storage, func(config *mapstructure.DecoderConfig) {
			config.TagName = "json"
		})
		if err != nil {
			log.Fatalf("读取数据库配置信息出错：%s", err)
		}
		err = viper.UnmarshalKey("cache", &redis, func(config *mapstructure.DecoderConfig) {
			config.TagName = "json"
		})
		if err != nil {
			log.Fatalf("读取数据库配置信息出错：%s", err)
		}
		// 监控JwtSecret
		viper.WatchConfig()
		viper.OnConfigChange(func(fsnotify.Event) {
			jwtSecret = viper.GetString("app.jwt_secret")
		})
	})
}

func GetJwtSecret() string {
	return jwtSecret
}

func GetHttpPort() int {
	return httpPort
}

func GetESUrl() string {
	return esUrl
}

func GetStorage() Storage {
	return storage
}

func GetCache() Redis {
	return redis
}
