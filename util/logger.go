package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var logOutPut zerolog.ConsoleWriter

var (
	pool sync.Pool
)

func init() {
	loggerFile := viper.GetString("logger.file_path")
	loggerlevel := viper.GetInt("logger.level")
	// 初始化日志配置
	zerolog.SetGlobalLevel(zerolog.Level(loggerlevel))
	if loggerFile == "" {
		logOutPut = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		err := os.MkdirAll(loggerFile[:strings.LastIndexByte(loggerFile, '/')+1], 0777)
		if err != nil {
			panic(fmt.Errorf("打开日志文件[%s]失败 \n", loggerFile))
		}
		file, err := os.Create(loggerFile)
		if err != nil {
			panic(fmt.Errorf("打开日志文件[%s]失败 \n", loggerFile))
		}
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		logOutPut = zerolog.ConsoleWriter{Out: io.MultiWriter(file, os.Stdout), TimeFormat: time.RFC3339}
	}
	logOutPut.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	logOutPut.FormatMessage = func(i interface{}) string {
		if i != nil {
			return fmt.Sprintf("***%s****", i)
		}
		return ""
	}
	logOutPut.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	logOutPut.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	pool = sync.Pool{New: func() interface{} {
		return zerolog.New(logOutPut).With().Timestamp().Logger()
	}}
}

func NewLogger() zerolog.Logger {
	return pool.Get().(zerolog.Logger)
}

func PutLogger(logger zerolog.Logger) {
	pool.Put(logger)
}
