package api

import (
	"gtpl/app/controller"

	"github.com/gin-gonic/gin"
	"github.com/soxft/gokt/gpi"
)

func initRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		gpi.New(c).Success("pong")
	})

	r.GET("/", controller.WelcomeIndex)

	// 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"success": false,
			"message": "route not found",
			"data":    nil,
		})
		gpi.New(c).FailWithHttpCode(404, "route not found")
	})
}
