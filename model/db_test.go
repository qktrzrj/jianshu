package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/unrotten/builder"
	"github.com/unrotten/hello-world-web/util"
	"github.com/unrotten/sqlex"
	"testing"
)

func Test_insertOne(t *testing.T) {
	tx, err := DB.Beginx()
	assert.NoError(t, err)
	logger := util.NewLogger()
	ctx := context.WithValue(context.Background(), "tx", tx)
	ctx = context.WithValue(ctx, "logger", logger)
	id, _ := idfetcher.NextID()
	result := insertOne(ctx, `"user"`, cv{
		"id":       int64(id),
		"username": "unrotten",
		"email":    "unrotten7@gmail.com",
		"password": "1",
		"avatar":   "",
		"gender":   "man",
		"state":    "normal",
		"root":     true,
	})
	assert.Equal(t, true, result.success)
	_ = tx.Commit()
}

func Test_update(t *testing.T) {
	tx, err := DB.Beginx()
	assert.NoError(t, err)
	logger := util.NewLogger()
	ctx := context.WithValue(context.Background(), "tx", tx)
	ctx = context.WithValue(ctx, "logger", logger)
	result := update(ctx, `"user"`, cv{"password": "2144"}, where{sqlex.Eq{"username": "unrotten"}})
	assert.Equal(t, true, result.success)
	_ = tx.Commit()
}

func Test_selectOne(t *testing.T) {
	tx, err := DB.Beginx()
	assert.NoError(t, err)
	logger := util.NewLogger()
	ctx := context.WithValue(context.Background(), "tx", tx)
	ctx = context.WithValue(ctx, "logger", logger)
	//r, err := tx.Exec(`SELECT "password" FROM "user" WHERE deleted_at is null AND username = $1 LIMIT 1`, "unrotten")
	//result := assertSqlResult(r, err, logger)
	result := selectOne(ctx, `"user"`, where{sqlex.Eq{"username": "unrotten"}}, `"password"`)
	assert.Equal(t, true, result.success)
	pwd, ok := builder.Get(result.b, "password")
	assert.Equal(t, true, ok)
	assert.Equal(t, "2144", pwd)
	_ = tx.Commit()
}

func Test_remove(t *testing.T) {
	tx, err := DB.Beginx()
	assert.NoError(t, err)
	logger := util.NewLogger()
	ctx := context.WithValue(context.Background(), "tx", tx)
	ctx = context.WithValue(ctx, "logger", logger)

	result := remove(ctx, `"user"`, where{sqlex.Eq{"username": "unrotten"}})
	assert.Equal(t, true, result.success)
	assert.Equal(t, false, selectOne(ctx, `"user"`, where{sqlex.Eq{"username": "unrotten"}}).success)
	assert.Equal(t, true, selectReal(ctx, `"user"`, where{sqlex.Eq{"username": "unrotten"}}).success)
	_ = tx.Commit()
}
