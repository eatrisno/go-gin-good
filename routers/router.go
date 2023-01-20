package routers

import (
	"github.com/eatrisno/go-gin-good/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", api.Ping)

	return r
}
