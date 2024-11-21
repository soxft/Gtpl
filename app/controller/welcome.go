package controller

import (
	"errors"
	"gtpl/app/model"
	"gtpl/process/db"

	"github.com/gin-gonic/gin"
	"github.com/soxft/gokt/gpi"
	"gorm.io/gorm"
)

func WelcomeIndex(c *gin.Context) {
	var counts model.Counter
	err := db.D.Model(&model.Counter{}).First(&counts).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.D.Create(&model.Counter{Count: 1})
		return
	}

	db.D.Model(&model.Counter{}).Where("1=1").Update("count", gorm.Expr("count + ?", 1))

	gpi.New(c).SuccessWithData("Welcome to soxft/Gtpl", gin.H{
		"count":    counts.Count,
		"updateAT": counts.UpdateAt,
	})
}
