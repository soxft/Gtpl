package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gtpl/app/model"
	"gtpl/process/db"
)

func WelcomeIndex(c *gin.Context) {
	counts := db.D.Model(&model.Counter{}).First(&model.Counter{})

	if errors.Is(counts.Error, gorm.ErrRecordNotFound) {
		db.D.Create(&model.Counter{Count: 1})
		return
	}

	db.D.Model(&model.Counter{}).Where("1=1").Update("count", gorm.Expr("count + ?", 1))

	c.JSON(200, gin.H{
		"success": true,
		"message": "Welcome to soxft/Gtpl",
		"data":    gin.H{"count": counts},
	})
	//
	//gpi.New(c).SuccessWithData("Welcome to soxft/Gtpl", gin.H{
	//	"count": counts,
	//})
}
