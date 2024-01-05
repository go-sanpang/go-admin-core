package loghelp

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	l *Logger
)

func Set(lo *Logger) {
	l = lo
}

func Get() *Logger {
	return l
}

func GetLogger(name, key string, config *LogConfig) *Logger {
	return l.Derive(name).SetLevel(getLevel(key, config)).AddCallerSkip(1)
}

func getLevel(key string, config *LogConfig) zapcore.Level {
	if _, exist := config.LogLevel.Specific[key]; exist {
		return config.LogLevel.Specific[key]
	}
	return config.LogLevel.Default
}

func Derive(s string) *Logger {
	var names []string
	if l.name == `` {
		names = append(names, s)
	} else {
		names = append(names, l.name, s)
	}

	return newLogger(l.underlying, strings.Join(names, "."), -1, true, l.fields...)
}

func With(fields ...zap.Field) *Logger {
	fields = append(l.fields, fields...)
	return newLogger(l.underlying.With(fields...), l.name, -1, false)
}

func Info(msg string, fields ...zap.Field) {
	l.underlying.Info(msg, fields...)
}

func Sync() error {
	return l.underlying.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	l.underlying.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	l.underlying.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.underlying.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	l.underlying.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	l.underlying.Panic(msg, fields...)
}

func SetLevel(level zapcore.Level) *Logger {
	var allCore []zapcore.Core
	if writeSyncer != nil {
		allCore = append(allCore, zapcore.NewCore(
			encoder,
			writeSyncer,
			level,
		))
	}

	for _, inputCore := range inputCores {
		if inputCore != nil {
			allCore = append(allCore, inputCore)
		}
	}

	allCore = append(allCore, zapcore.NewCore(encoder, os.Stdout, level))

	lg := zap.New(zapcore.NewTee(allCore...)).With(l.fields...)
	lg = lg.WithOptions(zap.AddCaller())

	return newLogger(lg, l.name, l.skip, true, l.fields...)
}

func Start(taskId string) *Logger {
	return l.With(zap.String(`任务ID`, taskId))
}

func AddCallerSkip(skip int) *Logger {
	return newLogger(l.underlying, l.name, skip, false)
}
