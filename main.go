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

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	idleTimeout := setting.ServerSetting.IdleTimeout
	readHeaderTimeout := setting.ServerSetting.ReadHeaderTimeout

	endPoint := fmt.Sprintf("%s:%d", setting.ServerSetting.HttpHost, setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:              endPoint,
		Handler:           routersInit,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	logging.Log.Info().Msg("Server is running...")

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logging.Log.Info().Msgf("caught sig: %+v", sig)
		logging.Log.Info().Msg("Wait for 3 second to finish processing")
		time.Sleep(3 * time.Second)
		if err := server.Shutdown(context.Background()); err != nil {
			logging.Log.Error().Msgf("Server Shutdown Failed:%+v", err)
		}
		logging.Log.Info().Msg("Server exiting")
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logging.Log.Info().Msg("Server closed under request")
		} else {
			logging.Log.Error().Msgf("Server closed unexpected %+v", err)
		}
	}
}
