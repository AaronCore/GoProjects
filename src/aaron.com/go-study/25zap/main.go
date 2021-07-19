package main

import (
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {
	//f1()
	//f2()
	r := gin.New()
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})
	r.Run(":8080")
}

func f1() {
	logger := zap.NewExample()
	defer logger.Sync()
	url := "https://gocn.zip"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func f2() {
	logger := zap.NewExample()
	defer logger.Sync()
	url := "https://gocn.vip"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
