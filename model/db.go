package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
	"github.com/unrotten/sqlex"
	"log"
	"time"
)

var (
	DB        *sqlx.DB
	PSql      sqlex.StatementBuilderType
	IdFetcher *sonyflake.Sonyflake
)

// 初始化数据库连接
func init() {
	// 获取数据库配置信息
	user := viper.Get("storage.user")
	password := viper.Get("storage.password")
	host := viper.Get("storage.host")
	port := viper.Get("storage.port")
	dbname := viper.Get("storage.dbname")

	// 连接数据库
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB = sqlx.MustOpen("postgres", psqlInfo)
	if err := DB.Ping(); err != nil {
		log.Fatalf("连接数据库失败:%s", err)
	}

	// 初始化sql构建器，指定format形式
	PSql = sqlex.StatementBuilder.PlaceholderFormat(sqlex.Dollar)

	// 初始化sonyflake
	st := sonyflake.Settings{
		StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	}
	IdFetcher = sonyflake.NewSonyflake(st)
}
