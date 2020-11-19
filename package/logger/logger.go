package logger

import (
	"ginn/config"
	errmsg "ginn/utils/code"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	FILENAME = "./log/ginn.log"
)

var Logger = NewLogger(config.Gconfig.Server.Mode)

func NewLogger(mode string) (sugarLogger *zap.SugaredLogger) {
	var err error
	writeSyncer := getLogWriter()
	var logger *zap.Logger
	switch mode {
	case gin.DebugMode:
		logger, err = zap.NewDevelopment(zap.AddCaller())
		errmsg.CheckMsgWithPanic(err)
	case gin.ReleaseMode:
		encoder := getEncoder(mode)
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
		logger = zap.New(core, zap.AddCaller())
	}
	sugarLogger = logger.Sugar()
	return
}

func getEncoder(mode string) zapcore.Encoder {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	switch mode {
	case gin.DebugMode:
		return zapcore.NewConsoleEncoder(encoder)
	case gin.ReleaseMode:
		return zapcore.NewJSONEncoder(encoder)
	default:
		return zapcore.NewJSONEncoder(encoder)
	}

}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   FILENAME,
		MaxSize:    200,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Sync() {
	_ = Logger.Sync()
}
