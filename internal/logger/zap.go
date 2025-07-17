package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	// Konfigurasi encoder (format log)
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
	// Buka file log
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// Set level log (di sini: info ke atas)
	logLevel := zapcore.InfoLevel
	core := zapcore.NewTee(
		// Output ke terminal (console)
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(os.Stdout), logLevel),
		// Output ke file
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.AddSync(logFile), logLevel),
	)

	// Buat logger dengan tambahan caller info
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}
