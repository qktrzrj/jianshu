package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	RunMode string

	HttpPort string

	MailHost string
	MailPort int
	MailAddr string
	MailPwd  string

	JwtSecret string
)

func init() {
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}

	// 设置默认配置
	viper.SetDefault("run_mode", "0")

	viper.SetDefault("http.port", "8008")

	viper.SetDefault("logger.level", "debug")

	viper.SetDefault("storage.user", "admin")
	viper.SetDefault("storage.password", "admin")
	viper.SetDefault("storage.host", "localhost")
	viper.SetDefault("storage.port", 5432)
	viper.SetDefault("storage.dbname", "postgres")

	// 获取配置信息
	RunMode = viper.GetString("run_mode")

	HttpPort = viper.GetString("http.port")

	MailHost = viper.GetString("mail.host")
	MailPort = viper.GetInt("mail.port")
	MailAddr = viper.GetString("mail.email")
	MailPwd = viper.GetString("mail.password")

	JwtSecret = viper.GetString("app.jwt_secret")
}
