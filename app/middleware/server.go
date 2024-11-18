package middleware

import "github.com/gin-gonic/gin"

func Server() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Server", "gokt")
		c.Next()
	}
}
