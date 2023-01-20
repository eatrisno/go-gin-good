package api

import (
	"net/http"

	"github.com/eatrisno/go-gin-good/resources/app"
	"github.com/eatrisno/go-gin-good/resources/e"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/hello [get]
func Helloworld(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "helloworld!",
	})
}
