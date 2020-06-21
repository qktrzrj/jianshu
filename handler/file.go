package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"io"
	"os"
)

func registerFileHandler(schema *schemabuilder.Schema) {
	schema.Query().FieldFunc("DownLoad", func(ctx context.Context, arg struct {
		FileName string `graphql:"fileName"`
	}) error {
		gctx := graphql.GetContext(ctx)
		logger := ctx.Value("logger").(zerolog.Logger)

		file, err := os.Open("/static/image/" + arg.FileName)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return errors.New("文件下载失败")
		}

		gctx.Writer.Header().Add("Content-Type", "application/octet-stream")
		gctx.Writer.Header().Add("content-disposition", "attachment; filename=\""+arg.FileName+"\"")
		_, err = io.Copy(gctx.Writer.ResponseWriter, file)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return errors.New("文件下载失败")
		}
		return nil
	})

	schema.Mutation().FieldFunc("Upload", func(ctx context.Context, arg struct {
		File schemabuilder.Upload `graphql:"file"`
	}) (string, error) {
		logger := ctx.Value("logger").(zerolog.Logger)

		file := arg.File.File
		defer file.Close()

		h := md5.New()
		io.Copy(h, file)
		fileName := hex.EncodeToString(h.Sum(nil))
		openFile, err := os.OpenFile("./static/image/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return "", errors.New("文件上传失败")
		}
		defer openFile.Close()
		file.Seek(0, 0)
		w, err := io.Copy(openFile, file)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return "", errors.New("文件上传失败")
		}
		if w != arg.File.Size {
			return "", errors.New("文件上传失败")
		}
		return fileName, nil
	}, middleware.LoginNeed())
}
