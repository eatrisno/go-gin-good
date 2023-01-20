package setting

import "time"

type App struct {
	JwtSecret string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	LogLevel     string
	HttpHost     string
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func Setup() {
	AppSetting.JwtSecret = "secret"
	AppSetting.LogSavePath = "logs/"
	AppSetting.LogSaveName = "log"
	AppSetting.LogFileExt = "log"
	AppSetting.TimeFormat = "20060102"

	ServerSetting.ReadTimeout = 60 * time.Second
	ServerSetting.WriteTimeout = 60 * time.Second

	ServerSetting.RunMode = "debug"

	ServerSetting.HttpHost = "localhost"
	ServerSetting.HttpPort = 8000
	ServerSetting.LogLevel = "DEBUG"
}
