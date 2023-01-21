package main

import (
	"fmt"
	"net/http"

	"github.com/eatrisno/go-gin-good/resources/logging"
	"github.com/eatrisno/go-gin-good/resources/setting"

	"github.com/eatrisno/go-gin-good/routers"
)

func init() {
	setting.Setup()
	logging.Info("App is starting...")

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

	logging.Info("Server is running...")

	server.ListenAndServe()
}
