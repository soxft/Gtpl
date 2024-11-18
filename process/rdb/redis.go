package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gtpl/library/zag"
	"time"
)

var R *redis.Client

func Init() {
	zag.L.Infof("Redis trying connect to tcp://%s/%d", viper.GetString("REDIS_ADDR"), viper.GetInt("REDIS_DB"))

	R = redis.NewClient(&redis.Options{
		Addr:            viper.GetString("REDIS_ADDR"),
		Password:        viper.GetString("REDIS_PWD"),
		DB:              viper.GetInt("REDIS_DB"),
		MinIdleConns:    viper.GetInt("REDIS_MIN_IDLE"),
		MaxIdleConns:    viper.GetInt("REDIS_MAX_IDLE"),
		MaxRetries:      viper.GetInt("REDIS_MAX_RETRIES"),
		ConnMaxLifetime: 5 * time.Minute,
		MaxActiveConns:  viper.GetInt("REDIS_MAX_ACTIVE"),
	})

	// test redis
	pong, err := R.Ping(context.Background()).Result()
	if err != nil {
		zag.L.Fatalf("redis ping failed: %v", err)
	}

	zag.L.Infof("redis init success, pong: %s ", pong)
}
