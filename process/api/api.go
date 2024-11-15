package api

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"gtpl/library/zag"

	"github.com/spf13/viper"
	"time"
)

// Run api service
func Run() {
	r := gin.New()

	r.Use(ginzap.Ginzap(zag.Z, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zag.Z, true))

	if viper.GetBool("debug") {
		zag.L.Info("api service run in debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	initRoutes(r)

	if err := r.Run(viper.GetString("api.listen")); err != nil {
		zag.L.Fatalf("api service run failed: %s", err)
	}
}
