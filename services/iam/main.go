package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandeep-jaiswar/engineering/pkg/config"
	"github.com/sandeep-jaiswar/engineering/pkg/db"
	"github.com/sandeep-jaiswar/engineering/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	err := logger.InitLogger("debug", "development")
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}
	defer logger.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Logger.Error("Failed to load configuration")
	}

	logger.Logger.Info("Config: ", zap.String("AppName", cfg.AppName), zap.String("IamPort", cfg.IamPort))

	client, err := db.Open(cfg.Database)
	if err != nil {
		logger.Logger.Panic("failed to connect to database: " + err.Error())
	}
	defer client.Close()

	if err := db.Migrate(client); err != nil {
		logger.Logger.Panic("failed to run migrations: " + err.Error())
	}

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		logger.Logger.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
		logger.Logger.Info("Handled root endpoint",
			zap.Int("status_code", http.StatusOK),
		)
	})

	if err := r.Run(":" + cfg.IamPort); err != nil {
		logger.Logger.Fatal("Failed to run server", zap.Error(err))
	}
}
