package main

import (
	shortUrl "Project/StartGoLang/Controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	shortUrl.RegisterRoutes(r)
	r.Run()
}
