package main

import (
	"fmt"
	"net/http"

	"github.com/eatrisno/go-gin-good/resources/setting"

	"github.com/eatrisno/go-gin-good/routers"
)

func init() {
	setting.Setup()
}
func main() {
	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf("%s:%d", setting.ServerSetting.HttpHost, setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
}
