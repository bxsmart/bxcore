package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
)

func Initialize(config ...zap.Config) *zap.Logger {
	var cfg zap.Config
	if len(config) <= 0 {
		var err error
		cfg = zap.NewDevelopmentConfig()
		if logger, err = cfg.Build(); err != nil {
			panic(err)
			return nil
		}
	} else {
		cfg = config[0]

		sink := lumberjack.Logger{
			Filename:   cfg.OutputPaths[0], // 日志文件路径
			MaxSize:    500,                // megabytes
			MaxBackups: 20,                 // 最多保留3个备份
			MaxAge:     30,                 //days
			Compress:   true,               // 是否压缩 disabled by default
		}

		errSink := lumberjack.Logger{
			Filename:   cfg.ErrorOutputPaths[0], // 日志文件路径
			MaxSize:    1024,                    // megabytes
			MaxBackups: 1,                       // 最多保留3个备份
			MaxAge:     30,                      //days
			Compress:   true,                    // 是否压缩 disabled by default
		}

		cfg.EncoderConfig = zap.NewProductionEncoderConfig()

		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding

		// 实现两个判断日志等级的interface
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.WarnLevel
		})

		warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel
		})

		// zapcore.NewTee 控制不同级别日志输出到不同源
		var zapcores []zapcore.Core
		encoder := zapcore.NewConsoleEncoder(cfg.EncoderConfig)
		zapcores = append(zapcores, zapcore.NewCore(encoder, zapcore.AddSync(&sink), infoLevel))
		zapcores = append(zapcores, zapcore.NewCore(encoder, zapcore.AddSync(&errSink), warnLevel))

		// 开发模式, debug级别以上全部打印到控制台
		if cfg.Development {
			zapcores = append(zapcores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
		}
		core := zapcore.NewTee(zapcores...)

		// 设置初始化字段 地址后面追加字段
		//filed := zap.Fields(zap.String("serviceName", "Starfire-OTC"))

		logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1), zap.ErrorOutput(zapcore.AddSync(&errSink)))
	}

	sugaredLogger = logger.Sugar()
	logger.Info(">>>> DefaultLogger INIT success <<<<")

	return logger
}

func GetLogger() *zap.Logger {
	return logger
}

func IsInit() bool {
	return nil != logger
}
