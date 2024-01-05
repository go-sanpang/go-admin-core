package loghelp

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志器的实现
type Logger struct {
	underlying *zap.Logger     // 底层日志器
	fields     []zapcore.Field // 嵌套的字段
	name       string          // 对应的名称
	skip       int             // skip
}

// newLogger 生成一个日志器
func newLogger(underlying *zap.Logger, name string, skip int, setName bool, fields ...zapcore.Field) *Logger {
	result := &Logger{underlying: underlying, name: name, fields: fields}

	if setName {
		result.underlying = result.underlying.Named(name)
	}

	if skip >= 0 {
		result.underlying = result.underlying.WithOptions(zap.AddCallerSkip(skip))
		result.skip = skip
	}

	return result
}

func (l *Logger) Derive(s string) *Logger {
	var names []string
	if l.name == `` {
		names = append(names, s)
	} else {
		names = append(names, l.name, s)
	}
	return newLogger(l.underlying, strings.Join(names, "."), -1, true, l.fields...)
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	fields = append(l.fields, fields...)
	return newLogger(l.underlying.With(fields...), l.name, -1, false)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.underlying.Info(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.underlying.Sync()
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.underlying.Debug(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.underlying.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.underlying.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.underlying.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.underlying.Panic(msg, fields...)
}

func (l *Logger) SetLevel(level zapcore.Level) *Logger {
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

func (l *Logger) Start(taskId string) *Logger {
	return l.With(zap.String(`任务ID`, taskId))
}

func (l *Logger) AddCallerSkip(skip int) *Logger {
	return newLogger(l.underlying, l.name, skip, false)
}
