package setting

import (
	"os"
	"regexp"
	"time"
)

type App struct {
	JwtSecret string

	SkipPath       []string
	SkipPathRegexp *regexp.Regexp
}

var AppSetting = &App{}

type Server struct {
	LogLevel          string
	HttpHost          string
	RunMode           string
	HttpPort          int
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	MaxHeaderBytes    int
}

var ServerSetting = &Server{}

func Setup() {
	AppSetting.JwtSecret = "secret"

	AppSetting.SkipPath = []string{"/ping"}
	AppSetting.SkipPathRegexp = regexp.MustCompile("^/[swagger].*.$")

	ServerSetting.ReadTimeout = 30 * time.Second
	ServerSetting.WriteTimeout = 30 * time.Second
	ServerSetting.IdleTimeout = 30 * time.Second
	ServerSetting.ReadHeaderTimeout = 1000 * time.Millisecond
	ServerSetting.MaxHeaderBytes = 1024

	ServerSetting.HttpHost = "localhost"
	ServerSetting.HttpPort = 8000
	ServerSetting.LogLevel = "INFO"

	if os.Getenv("ENV") == "production" {
		ServerSetting.RunMode = "release"
	} else {
		ServerSetting.RunMode = "debug"
	}

}
