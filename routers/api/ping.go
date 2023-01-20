package api

import (
	"net/http"

	"github.com/eatrisno/go-gin-good/resources/app"
	"github.com/eatrisno/go-gin-good/resources/e"
	"github.com/gin-gonic/gin"
)

// @Summary	Get a ping response
// @Produce	json
// @Success	200	{object}	app.Response
// @Failure	500	{object}	app.Response
// @Router		/ping [get]
func Ping(c *gin.Context) {

	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "pong",
	})

}
