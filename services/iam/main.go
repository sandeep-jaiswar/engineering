package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/engineering/common/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}
	defer logger.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration")
	}

	logger.Info("Config: ",zap.String("AppName", cfg.AppName),zap.String("IamPort", cfg.IamPort))

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		logger.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
		logger.Info("Handled root endpoint",
			zap.Int("status_code", http.StatusOK),
		)
	})

	if err := r.Run(":"+cfg.IamPort); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
