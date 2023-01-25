package setting

import (
	"os"
	"regexp"
	"time"

	"github.com/rs/zerolog"
)

type App struct {
	JwtSecret string

	SkipPath       []string
	SkipPathRegexp *regexp.Regexp

	BeautifulLogging bool
}

var AppSetting = &App{}

type Server struct {
	LogLevel          zerolog.Level
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

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	AppSetting.JwtSecret = "secret"

	AppSetting.SkipPath = []string{"/ping"}
	AppSetting.SkipPathRegexp = regexp.MustCompile("^/[swagger].*.$")

	ServerSetting.ReadTimeout = 30 * time.Second
	ServerSetting.WriteTimeout = 30 * time.Second
	ServerSetting.IdleTimeout = 30 * time.Second
	ServerSetting.ReadHeaderTimeout = 100 * time.Millisecond
	ServerSetting.MaxHeaderBytes = 1024

	ServerSetting.HttpHost = "localhost"
	ServerSetting.HttpPort = 8000

	if os.Getenv("ENV") == "production" {
		ServerSetting.RunMode = "release"
		ServerSetting.LogLevel = zerolog.WarnLevel
	} else {
		ServerSetting.RunMode = "debug"
		ServerSetting.LogLevel = zerolog.DebugLevel
	}

}
