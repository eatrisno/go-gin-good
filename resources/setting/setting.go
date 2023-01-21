package setting

import (
	"os"
	"regexp"
	"time"
)

type App struct {
	JwtSecret string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

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
}

var ServerSetting = &Server{}

func Setup() {
	AppSetting.JwtSecret = "secret"
	AppSetting.LogSavePath = "logs/"
	AppSetting.LogSaveName = "log"
	AppSetting.LogFileExt = "log"
	AppSetting.TimeFormat = "20060102"

	AppSetting.SkipPath = []string{"/ping"}
	AppSetting.SkipPathRegexp = regexp.MustCompile("^/swagger/")

	ServerSetting.ReadTimeout = 60 * time.Second
	ServerSetting.WriteTimeout = 60 * time.Second
	ServerSetting.IdleTimeout = 30 * time.Second
	ServerSetting.ReadHeaderTimeout = 1000 * time.Millisecond
	if os.Getenv("ENV") == "production" {
		ServerSetting.RunMode = "release"
	} else {
		ServerSetting.RunMode = "debug"
	}

	ServerSetting.HttpHost = "localhost"
	ServerSetting.HttpPort = 8000
	ServerSetting.LogLevel = "INFO"
}
