package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
)

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Name",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "St",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/app.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // 圧縮するか
		LocalTime:  true, // バックアップファイル名の時間
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // ログフォーマット
		w,                                        // 出力先 lumberjack
		//zap.ErrorLevel,                           // ログレベル INFO
		zap.InfoLevel, // ログレベル INFO
	)
	logger = zap.New(core)
}

// GetLogger 初期化済みのLoggerを返却する
func GetLogger() *zap.Logger {
	return logger
}
