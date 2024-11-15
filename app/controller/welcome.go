package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/soxft/gokt/gpi"
)

func WelcomeIndex(c *gin.Context) {
	gpi.New(c).Success("Welcome to soxft/Gtpl")
}
