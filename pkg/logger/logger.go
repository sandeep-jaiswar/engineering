package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func InitLogger(logLevel string, environment string) error {
	var cfg zap.Config

	if environment == "development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	level := zapcore.InfoLevel
	if err := level.Set(logLevel); err == nil {
		cfg.Level.SetLevel(level)
	} else {
		return err
	}

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		return err
	}

	Sugar = Logger.Sugar()
	zap.ReplaceGlobals(Logger)
	return nil
}

func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}