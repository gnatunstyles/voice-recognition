package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnatunstyles/voice-recognition/internal/routes"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)

	r.Run()
}
