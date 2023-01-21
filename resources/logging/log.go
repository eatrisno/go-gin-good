package logging

import (
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type config struct {
	utc            bool
	skipPath       []string
	skipPathRegexp *regexp.Regexp
}

var cfg = &config{
	utc:            false,
	skipPath:       []string{"/ping"},
	skipPathRegexp: regexp.MustCompile("/swagger/*"),
}

var Log = zerolog.New(os.Stdout)

func init() {
	if os.Getenv("ENV") != "production" {
		Log = Log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Caller().Logger()
	}
}

func Middleware() gin.HandlerFunc {
	var skip map[string]struct{}
	if length := len(cfg.skipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range cfg.skipPath {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()
		track := true

		if _, ok := skip[path]; ok {
			track = false
		}

		if track && cfg.skipPathRegexp != nil && cfg.skipPathRegexp.MatchString(path) {
			track = false
		}

		if track {
			end := time.Now()
			if cfg.utc {
				end = end.UTC()
			}
			latency := end.Sub(start)

			if latency > time.Minute {
				latency = latency.Truncate(time.Second)
			}

			l := Log.With().
				Str("client_id", c.ClientIP()).
				Str("method", c.Request.Method).
				Int("status_code", c.Writer.Status()).
				Int("body_size", c.Writer.Size()).
				Str("path", c.Request.URL.Path).
				Dur("latency", latency).
				Logger()

			msg := "Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}
			switch {
			case c.Writer.Status() >= http.StatusInternalServerError:
				l.WithLevel(zerolog.ErrorLevel).Msg(msg)
			case c.Writer.Status() >= http.StatusBadRequest:
				l.WithLevel(zerolog.WarnLevel).Msg(msg)
			default:
				l.WithLevel(zerolog.InfoLevel).Msg(msg)
			}
		}
	}
}
