package util

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql/context"
	"github.com/shyptr/hello-world-web/setting"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var logOutPut zerolog.ConsoleWriter

var (
	pool sync.Pool
)

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.Level(setting.GetLoggerConfig().GetLevel()))
	if !setting.GetLoggerConfig().GetEnable() {
		logOutPut = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		err := os.MkdirAll(setting.GetLoggerConfig().GetPath()[:strings.LastIndexByte(setting.GetLoggerConfig().GetPath(), '/')+1], 0777)
		if err != nil {
			panic(fmt.Errorf("open logger file [%s] failed \n", setting.GetLoggerConfig().GetPath()))
		}
		file, err := os.Create(setting.GetLoggerConfig().GetPath())
		if err != nil {
			panic(fmt.Errorf("open logger file [%s] failed \n", setting.GetLoggerConfig().GetPath()))
		}
		context.SetLogger(log.New(io.MultiWriter(file, os.Stdout), log.Prefix(), log.Flags()))
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
