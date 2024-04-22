package ShortenUrlController

import (
	service "Project/StartGoLang/Service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/shortUrl/:shortUrl", service.ShortUrl)
	r.POST("/shorten", service.Shorten)
}
