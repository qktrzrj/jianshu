package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro/v2"
	"github.com/shyptr/hello-world-web/setting"
	"github.com/shyptr/hello-world-web/util"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"github.com/sony/sonyflake"
	"log"
	"time"
)

var (
	DB        *sqlog.DB
	PSql      sqlex.StatementBuilderType
	idfetcher *sonyflake.Sonyflake
)

func Init(*micro.Options) {

	DB = &sqlog.DB{Runner: sqlx.MustOpen(setting.GetStorageConfig().GetType(), setting.GetStorageConfig().GetURL())}
	if err := DB.Runner.(*sqlx.DB).Ping(); err != nil {
		log.Fatalf("connection db failed:%s", err)
	}
	DB.Logger = util.NewSQLogger(util.NewLogger())
	if setting.GetStorageConfig().GetType() == "postgres" {
		PSql = sqlex.StatementBuilder.PlaceholderFormat(sqlex.Dollar)
	}

	st := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return 1, nil
		},
		StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	}
	idfetcher = sonyflake.NewSonyflake(st)
}
