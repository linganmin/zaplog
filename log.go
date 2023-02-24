package zaplog

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"zaplog/metadatax"
)

// setLogWriter
func setLogWriter() zapcore.WriteSyncer {
	wd, _ := os.Getwd()
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s/%s.log", wd, "logs", time.Now().Format("2006-01-02")),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(io.MultiWriter(os.Stdout, lumberJackLogger))
}

func setEncoder() zapcore.Encoder {

	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "path_line",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return initLogger().Sugar()
	}

	var fields []zap.Field
	for _, key := range metadatax.Keys {
		if value, ok := ctx.Value(key).(string); ok {
			fields = append(fields, zap.String(key, value))
		}
	}

	return initLogger().With(fields...).Sugar()
}

func initLogger() *zap.Logger {
	core := zapcore.NewCore(setEncoder(), setLogWriter(), zapcore.DebugLevel)

	// 1.log filepath，2.error level print stack info
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel), zap.Development())

}
