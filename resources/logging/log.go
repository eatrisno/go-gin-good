package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func DefaultStructuredLogger() gin.HandlerFunc {
	return StructuredLogger(&log.Logger)
}

func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
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
