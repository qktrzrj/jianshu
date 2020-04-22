package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/setting"
	"github.com/shyptr/jianshu/util"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"github.com/sony/sonyflake"
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
	DB        *sqlog.DB
	PSql      sqlex.StatementBuilderType
	IdFetcher *sonyflake.Sonyflake
	once      sync.Once
)

func Init() {
	once.Do(func() {
		s := setting.GetStorage()
		pi := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			s.GetHost(), s.GetPort(), s.GetUsername(), s.GetPassword(), s.GetDBName())
		db := sqlx.MustOpen("postgres", pi)
		DB = &sqlog.DB{Runner: db, Logger: sqloger{logger: util.GetLogger()}}
		PSql = sqlex.StatementBuilder.PlaceholderFormat(sqlex.Dollar)
		IdFetcher = sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local),
			MachineID: func() (uint16, error) {
				return 0, nil
			},
		})
	})
}
