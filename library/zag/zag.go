package zag

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var L *zap.SugaredLogger
var Z *zap.Logger

// Init zap instance
func Init() {
	if viper.GetBool("DEBUG") {
		Z, _ = zap.NewDevelopment()

		Z.Sugar().Info("Running in debug mode")
	} else {
		// 不提示 debug
		Z, _ = zap.NewProduction()
	}

	L = Z.Sugar()
}
