package db

import (
	"fmt"
	"gtpl/app/model"
	"gtpl/library/zag"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var D *gorm.DB

func Init() {
	zag.L.Infof("mysql trying connect to tcp://%s:%s/%s", viper.GetString("MYSQL_USERNAME"), viper.GetString("MYSQL_ADDR"), viper.GetString("MYSQL_DATABASE"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("MYSQL_USERNAME"), viper.GetString("MYSQL_PASSWORD"), viper.GetString("MYSQL_ADDR"), viper.GetString("MYSQL_DATABASE"), viper.GetString("MYSQL_CHARSET"))

	logger := zapgorm2.New(zag.Z)
	logger.SetAsDefault()

	var err error
	D, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		zag.L.Fatalf("mysql error: %v", err)
	}

	sqlDb, err := D.DB()
	if err != nil {
		zag.L.Fatalf("mysql get db error: %v", err)
	}

	sqlDb.SetMaxOpenConns(viper.GetInt("MYSQL_MAX_OPEN"))
	sqlDb.SetMaxIdleConns(viper.GetInt("MYSQL_MAX_IDLE"))
	sqlDb.SetConnMaxLifetime(time.Duration(viper.GetInt("MYSQL_MAX_LIFETIME")) * time.Second)
	if err := sqlDb.Ping(); err != nil {
		zag.L.Fatalf("mysql connect error: %v", err)
	}

	// 自动同步表结构
	if err := D.AutoMigrate(model.Counter{}); err != nil {
		zag.L.Fatalf("mysql migrate error: %v", err)
	}

	zag.L.Infof("mysql init success")
}
