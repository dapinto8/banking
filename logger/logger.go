package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func config() (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.NewProductionConfig()
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

func Debug(message string, fields ...interface{}) {
	sugar.Debugw(message, fields...)
}

func Error(message string, fields ...interface{}) {
	sugar.Errorw(message, fields...)
}
