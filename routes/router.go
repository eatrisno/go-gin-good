package routes

import (
	"sync"

	_ "github.com/eatrisno/go-gin-good/docs"
	"github.com/eatrisno/go-gin-good/resources/logging"
	"github.com/eatrisno/go-gin-good/resources/setting"
	"github.com/eatrisno/go-gin-good/routes/api"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var routerPool = sync.Pool{
	New: func() interface{} {
		return gin.New()
	},
}

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	r := routerPool.Get().(*gin.Engine)
	r.Use(logging.Middleware())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", api.Ping)
	r.GET("/hello", api.Helloworld)

	return r
}
