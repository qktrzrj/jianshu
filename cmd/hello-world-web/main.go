package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrotten/hello-world-web/controller"
	"github.com/unrotten/hello-world-web/setting"
	"github.com/unrotten/hello-world-web/util"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	gin.SetMode(setting.RunMode)
	engine := gin.New()
	controller.Register(engine)

	addr := ":" + setting.HttpPort
	server := &http.Server{
		Addr:           addr,
		Handler:        engine,
		MaxHeaderBytes: 1 << 20,
	}

	logger := util.NewLogger()

	go func() {
		logger.Info().Msg(fmt.Sprintf("server run on:%s", addr))
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal().Caller().Err(err).Msg("server err")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info().Msg("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Caller().Err(err).Msg("Server Shutdown")
	}
	logger.Info().Msg("Server exiting")
}
