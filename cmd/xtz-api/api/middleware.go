package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf(`"%s %s" received`, c.Request.Method, c.Request.URL.Path)

		c.Next()

		statusCode := c.Writer.Status()
		log.Printf(`"%s %s" handled - %3d`, c.Request.Method, c.Request.URL.Path, statusCode)
	}
}
