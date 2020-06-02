package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/setting"
	"github.com/shyptr/jianshu/util"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"sync"
	"time"
)

type sqloger struct {
	logger zerolog.Logger
}

func (s sqloger) BeforeInFoLog(query string, args ...interface{}) {}

func (s sqloger) AfterInFoLog(executeTime time.Duration, query string, args ...interface{}) {
	s.logger.Info().Dur("executeTime", executeTime).Interface("[SQL]", fmt.Sprintf("%v : %v", query, args)).Send()
}

func (s sqloger) Async() bool {
	return false
}

func (s sqloger) Show() bool {
	return true
}

var (
	DB   *sqlog.DB
	PSql sqlex.StatementBuilderType
	once sync.Once
)

func Init() {
	once.Do(func() {
		s := setting.GetStorage()
		pi := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
			s.GetUsername(), s.GetPassword(), s.GetHost(), s.GetPort(), s.GetDBName())
		db := sqlx.MustOpen("mysql", pi)
		DB = &sqlog.DB{Runner: db, Logger: sqloger{logger: util.GetLogger()}}
		PSql = sqlex.StatementBuilder
	})
}
