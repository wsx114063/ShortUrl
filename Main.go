package main

import (
	shortUrl "Project/StartGoLang/Controller"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	shortUrl.RegisterRoutes(r)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
