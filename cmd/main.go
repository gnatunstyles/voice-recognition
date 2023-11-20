package main

import (
	"example.com/m/v2/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)

	r.Run()
}
