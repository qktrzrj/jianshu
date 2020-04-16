package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"time"
)

type Tag struct {
	Id        int64     `graphql:"id" db:"id"`
	Name      string    `graphql:"name" db:"name"`
	CreatedAt time.Time `graphql:"-" db:"created_at"`
	UpdatedAt time.Time `graphql:"-" db:"updated_at"`
	DeletedAt time.Time `graphql:"-" db:"deleted_at"`
}

func InsertTag(ctx context.Context, name string) error {
	id, err := idfetcher.NextID()
	if err != nil {
		return err
	}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("tag").Columns("id,name").Values(id, name).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("create new tag failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("create new tag failed")
	}
	return nil
}

func GetTags(ctx context.Context, name string) ([]string, error) {
	var tags []string
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("tag").Columns("name").RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, errors.New("fetch tags name list failed")
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, errors.New("fetch tags name list failed")
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
