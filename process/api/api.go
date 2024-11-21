package api

import (
	"gtpl/app/middleware"
	"gtpl/library/zag"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"time"

	"github.com/spf13/viper"
)

// Run api service
func Run() {
	if viper.GetBool("DEBUG") {
		zag.L.Info("api service run in debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(ginzap.Ginzap(zag.Z, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zag.Z, true))
	r.Use(middleware.Server())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	initRoutes(r)

	zag.L.Infof("api service run on %s", viper.GetString("API_LISTEN"))

	if err := r.Run(viper.GetString("API_LISTEN")); err != nil {
		zag.L.Fatalf("api service run failed: %s", err)
	}
}
