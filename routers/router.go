package routers

import (
	"github.com/eatrisno/go-gin-good/resources/logging"
	"github.com/eatrisno/go-gin-good/resources/setting"
	"github.com/eatrisno/go-gin-good/routers/api"

	_ "github.com/eatrisno/go-gin-good/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	r := gin.New()
	r.Use(logging.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", api.Ping)
	r.GET("/hello", api.Helloworld)

	return r
}
