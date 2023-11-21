package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"one_ser.com/config"
)

// encoder编码，输出位置（终端/文件），日志级别（控制哪些输出，哪些不输出）
var Log *zap.Logger

// 初始化log
func Init(cfg *config.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		//开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentConfig().EncoderConfig) //终端
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {

		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	// 如何? 日志默认输出app.log,错误日志再在app.err.log中记录一份
	Log = zap.New(core, zap.AddCaller()) // AddCaller添加调用栈信息

	zap.ReplaceGlobals(Log) // 替换全局zap包全局的logger
	zap.L().Info("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{ // 日志切分
		Filename:  filename,
		MaxSize:   maxSize,
		MaxBackups: maxBackup,
		MaxAge:    maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
