package main

import (
	"context"
	"fmt"
	"github.com/shyptr/graphql"
	"github.com/shyptr/jianshu/handler"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/setting"
	"github.com/shyptr/jianshu/util"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	setting.Init()
	util.InitLog()
	model.Init()
}

func main() {
	graphql.Use(middleware.CORS(), middleware.Logger(), middleware.Tx(model.DB), middleware.Recovery())
	mux := http.DefaultServeMux
	handler.Register(mux)

	addr := fmt.Sprintf(":%d", setting.GetHttpPort())
	server := http.Server{
		Addr:           addr,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}

	logger := util.GetLogger()
	go func() {
		logger.Info().Msgf("server run on:%s", addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal().AnErr("Server Error", err).Send()
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	logger.Info().Msg("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Caller().Err(err).Msg("Server Shutdown")
	}
	logger.Info().Msg("Server exiting")
}
