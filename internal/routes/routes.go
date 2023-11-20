package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gnatunstyles/voice-recognition/internal/handlers"
)

func InitRoutes(app *gin.Engine) {
	app.POST("/voice_cmd", handlers.TranscribeHandler)
}
