package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
)

func InitializeTest() *zap.Logger {
	var err error
	cfg := zap.NewDevelopmentConfig()
	logger, err = cfg.Build()
	if err != nil {
		panic(err)
		return nil
	}

	sugaredLogger = logger.Sugar()
	logger.Info("DefaultLogger INIT success")

	return logger
}

func Initialize(cfg zap.Config) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   "./logs", // 日志文件路径
		MaxSize:    1024,     // megabytes
		MaxBackups: 7,        // 最多保留3个备份
		MaxAge:     30,        //days
		Compress:   true,     // 是否压缩 disabled by default
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(&hook),
		cfg.Level,
	)
	logger := zap.New(core)

	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic(err)
		return nil
	}

	sugaredLogger = logger.Sugar()
	logger.Info("DefaultLogger INIT success")

	return logger
}

func IsInit() bool {
	return nil != logger
}
