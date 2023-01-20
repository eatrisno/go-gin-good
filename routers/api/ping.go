package api

import (
	"net/http"

	"github.com/eatrisno/go-gin-good/resources/app"
	"github.com/eatrisno/go-gin-good/resources/e"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "pong",
	})

}
