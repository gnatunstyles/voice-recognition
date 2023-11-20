package routes

import (
	"example.com/m/v2/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	app.POST("/voice_cmd", handlers.TranscribeHandler)
}
