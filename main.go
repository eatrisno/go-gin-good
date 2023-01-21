package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eatrisno/go-gin-good/resources/logging"
	"github.com/eatrisno/go-gin-good/resources/setting"

	"github.com/eatrisno/go-gin-good/routers"
)

func init() {
	setting.Setup()
	logging.Log.Info().Msg("App is starting...")
}

func main() {
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf("%s:%d", setting.ServerSetting.HttpHost, setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:              endPoint,
		Handler:           routersInit,
		ReadTimeout:       setting.ServerSetting.ReadTimeout,
		WriteTimeout:      setting.ServerSetting.WriteTimeout,
		IdleTimeout:       setting.ServerSetting.IdleTimeout,
		ReadHeaderTimeout: setting.ServerSetting.ReadHeaderTimeout,
		MaxHeaderBytes:    setting.ServerSetting.MaxHeaderBytes,
	}

	logging.Log.Info().Msg("Server is running...")

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		logging.Log.Info().Msgf("caught sig: %+v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logging.Log.Error().Msgf("Server Shutdown Failed:%+v", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logging.Log.Info().Msg("Server closed under request")
		} else {
			logging.Log.Error().Msgf("Server closed unexpected %+v", err)
		}
	}
}
