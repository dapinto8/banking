package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func config() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	return config.Build(zap.AddCallerSkip(1))
}

func init() {

	logger, err := config()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}

	sugar = logger.Sugar()
}

func Info(message string, fields ...interface{}) {
	sugar.Infow(message, fields...)
}
