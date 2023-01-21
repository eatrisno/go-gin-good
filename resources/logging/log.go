package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = log.Logger

func init() {
	if os.Getenv("ENV") != "production" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		param := gin.LogFormatterParams{
			TimeStamp:    time.Now(),
			Latency:      time.Since(start),
			ClientIP:     c.Request.RemoteAddr,
			Method:       c.Request.Method,
			StatusCode:   c.Writer.Status(),
			ErrorMessage: c.Errors.ByType(gin.ErrorTypePrivate).String(),
			BodySize:     c.Writer.Size(),
			Path:         c.Request.URL.String(),
		}
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		level := zerolog.TraceLevel
		switch {
		case param.StatusCode >= 200 && param.StatusCode < 300:
			level = zerolog.InfoLevel
		case param.StatusCode >= 300 && param.StatusCode < 400:
			level = zerolog.WarnLevel
		case param.StatusCode >= 400:
			level = zerolog.ErrorLevel
		}

		logger.WithLevel(level).
			Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Dur("latency", param.Latency).
			Msg(param.ErrorMessage)
	}
}

func Debug(message ...interface{}) {
	logger.Debug().Msg(fmt.Sprint(message...))
}

func Info(message ...interface{}) {
	logger.Info().Msg(fmt.Sprint(message...))
}

func Warn(message ...interface{}) {
	logger.Warn().Msg(fmt.Sprint(message...))
}

func Fatal(message ...interface{}) {
	logger.Fatal().Msg(fmt.Sprint(message...))
}

func Panic(message ...interface{}) {
	logger.Panic().Msg(fmt.Sprint(message...))
}
