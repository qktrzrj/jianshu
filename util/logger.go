package util

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	logPool sync.Pool
	once    sync.Once
)

func InitLog() {
	once.Do(func() {
		logOutPut := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
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
			return fmt.Sprintf("%s |", i)
		}

		logPool.New = func() interface{} {
			return zerolog.New(logOutPut).With().Timestamp().Logger()
		}
	})
}

func GetLogger() zerolog.Logger {
	logger := logPool.Get()
	if logger == nil {
		InitLog()
	}
	logger = logPool.Get()
	return logger.(zerolog.Logger)
}

func PutLogger(logger zerolog.Logger) {
	logPool.Put(logger)
}
