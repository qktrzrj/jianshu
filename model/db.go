package model

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
	"github.com/unrotten/builder"
	"github.com/unrotten/sqlex"
	"log"
	"os"
	"reflect"
	"time"
)

var (
	DB        *sqlx.DB
	psql      sqlex.StatementBuilderType
	idfetcher *sonyflake.Sonyflake
)

const defaultSkip int = 2

type cv map[string]interface{}

type where []sqlex.Sqlex

type result struct {
	b       builder.Builder
	success bool
}

// 初始化数据库连接
func init() {
	viper.AddConfigPath("../config") // 测试使用
	viper.ReadInConfig()
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
	psql = sqlex.StatementBuilder.PlaceholderFormat(sqlex.Dollar)
	sqlex.SetLogger(os.Stdout)

	// 初始化sonyflake
	st := sonyflake.Settings{
		StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	}
	idfetcher = sonyflake.NewSonyflake(st)
}

func get(query *sql.Rows, columnTypes []*sql.ColumnType, logger zerolog.Logger) result {
	dest := make([]interface{}, len(columnTypes))
	for index, col := range columnTypes {
		switch col.ScanType().String() {
		case "string", "interface {}":
			dest[index] = &sql.NullString{}
		case "bool":
			dest[index] = &sql.NullBool{}
		case "float64":
			dest[index] = &sql.NullFloat64{}
		case "int32":
			dest[index] = &sql.NullInt32{}
		case "int64":
			dest[index] = &sql.NullInt64{}
		case "time.Time":
			dest[index] = &sql.NullTime{}
		default:
			dest[index] = reflect.New(col.ScanType()).Interface()
		}
	}
	err := query.Scan(dest...)
	if err != nil {
		logger.Error().Caller(2).Err(err).Send()
		return result{success: false}
	}
	build := builder.EmptyBuilder
	for index, col := range columnTypes {
		switch val := dest[index].(type) {
		case driver.Valuer:
			var value interface{}
			switch col.ScanType().String() {
			case "string", "interface {}":
				value = dest[index].(*sql.NullString).String
			case "bool":
				value = dest[index].(*sql.NullBool).Bool
			case "float64":
				value = dest[index].(*sql.NullFloat64).Float64
			case "int32":
				value = dest[index].(*sql.NullInt32).Int32
			case "int64":
				value = dest[index].(*sql.NullInt64).Int64
			case "time.Time":
				value = dest[index].(*sql.NullTime).Time
			}
			build = builder.Set(build, col.Name(), value).(builder.Builder)
		default:
			build = builder.Set(build, col.Name(), val).(builder.Builder)
		}
	}
	return result{success: true, b: build}
}

func selectList(ctx context.Context, table string, where where, columns ...string) result {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlx.Tx)

	var selectBuilder sqlex.SelectBuilder
	if len(columns) > 0 {
		selectBuilder = psql.Select(columns...).From(table).Where("deleted_at is null")
	} else {
		selectBuilder = psql.Select("*").From(table).Where("deleted_at is null")
	}
	for _, arg := range where {
		selectBuilder = selectBuilder.Where(arg)
	}
	query, err := selectBuilder.RunWith(tx).Query()
	if err != nil {
		logger.Error().Caller(1).Err(err).Send()
		return result{success: false}
	}

	columnTypes, err := query.ColumnTypes()
	if err != nil {
		logger.Error().Caller(1).Err(err).Send()
		return result{success: false}
	}
	var resultSlice []interface{}
	for query.Next() {
		r := get(query, columnTypes, logger)
		if !r.success {
			return r
		}
		resultSlice = append(resultSlice, r.b)
	}
	return result{success: true, b: builder.Set(builder.EmptyBuilder, "list", resultSlice).(builder.Builder)}
}

func selectOne(ctx context.Context, table string, where where, columns ...string) result {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlx.Tx)

	var selectBuilder sqlex.SelectBuilder
	if len(columns) > 0 {
		selectBuilder = psql.Select(columns...).From(table).Where("deleted_at is null").Limit(1)
	} else {
		selectBuilder = psql.Select("*").From(table).Where("deleted_at is null").Limit(1)
	}
	for _, arg := range where {
		selectBuilder = selectBuilder.Where(arg)
	}
	query, err := selectBuilder.RunWith(tx).Query()
	if err != nil {
		logger.Error().Caller(1).Err(err).Send()
		return result{success: false}
	}

	columnTypes, err := query.ColumnTypes()
	if err != nil {
		logger.Error().Caller(1).Err(err).Send()
		return result{success: false}
	}

	if query.Next() {
		return get(query, columnTypes, logger)
	}
	return result{success: false}
}

func selectReal(ctx context.Context, table string, where where, columns ...string) result {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlx.Tx)

	var selectBuilder sqlex.SelectBuilder
	if len(columns) > 0 {
		selectBuilder = psql.Select(columns...).From(table).Where("deleted_at is not null")
	} else {
		selectBuilder = psql.Select("*").From(table).Where("deleted_at is not null")
	}
	for _, arg := range where {
		selectBuilder = selectBuilder.Where(arg)
	}
	query, err := selectBuilder.RunWith(tx).Query()
	if err != nil {
		logger.Error().Caller(1).Err(err).Send()
		return result{success: false}
	}

	columnTypes, err := query.ColumnTypes()
	if err != nil {
		logger.Error().Caller(1).Err(err).Send()
		return result{success: false}
	}
	var resultSlice []interface{}
	for query.Next() {
		r := get(query, columnTypes, logger)
		if !r.success {
			return r
		}
		resultSlice = append(resultSlice, r.b)
	}
	return result{success: true, b: builder.Set(builder.EmptyBuilder, "list", resultSlice).(builder.Builder)}
}

func insertOne(ctx context.Context, table string, cv cv) result {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlx.Tx)
	build := builder.EmptyBuilder
	cv["created_at"], cv["updated_at"] = time.Now(), time.Now()
	columns, values := make([]string, 0, len(cv)), make([]interface{}, 0, len(cv))
	for col, value := range cv {
		build = builder.Set(build, col, value).(builder.Builder)
		columns, values = append(columns, col), append(values, value)
	}
	r, err := psql.Insert(table).Columns(columns...).Values(values...).RunWith(tx).Exec()
	return assertSqlResult(r, err, logger)
}

func update(ctx context.Context, table string, cv cv, where where, directSet ...string) result {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlx.Tx)
	cv["updated_at"] = time.Now()
	updateBuilder := psql.Update(table).SetMap(cv).Where("deleted_at is null")
	for _, set := range directSet {
		updateBuilder = updateBuilder.DirectSet(set)
	}
	for _, arg := range where {
		updateBuilder = updateBuilder.Where(arg)
	}
	r, err := updateBuilder.RunWith(tx).Exec()
	return assertSqlResult(r, err, logger)
}

// note: if where is null,then will delete the whole table
func remove(ctx context.Context, table string, where where) result {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlx.Tx)

	updateBuilder := psql.Update(table).Set("deleted_at", time.Now()).Where("deleted_at is null")
	for _, arg := range where {
		updateBuilder = updateBuilder.Where(arg)
	}
	r, err := updateBuilder.RunWith(tx).Exec()
	return assertSqlResult(r, err, logger)
}

func assertSqlResult(r sql.Result, err error, logger zerolog.Logger, skip ...int) result {
	sk := defaultSkip
	if len(skip) > 0 {
		sk += skip[0]
	}
	if err != nil {
		logger.Error().Caller(sk).Err(err).Send()
		return result{success: false}
	}
	affected, err := r.RowsAffected()
	if err != nil {
		logger.Error().Caller(2).Err(err).Send()
		return result{success: false}
	}
	if affected == 0 {
		return result{success: false}
	}
	return result{success: true}
}
