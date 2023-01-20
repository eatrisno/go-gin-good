package routers

import (
	_ "github.com/eatrisno/go-gin-good/docs"
	"github.com/eatrisno/go-gin-good/routers/api"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", api.Ping)
	r.GET("/hello", api.Helloworld)

	return r
}
