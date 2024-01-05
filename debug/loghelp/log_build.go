package loghelp

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

var (
	encoder     zapcore.Encoder
	writeSyncer zapcore.WriteSyncer
	inputCores  []zapcore.Core
)

// Build 构建日志器
func Build(lConfig *LogConfig, cores ...zapcore.Core) (logger *Logger) {
	var (
		err              error
		underlyingLogger *zap.Logger
		allCores         []zapcore.Core
	)

	inputCores = cores

	cfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.ErrorLevel),
		Development:      true,
		Encoding:         "console",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	location, err := time.LoadLocation(lConfig.TimeZone)
	if err != nil {
		panic(fmt.Sprintf(`加载时区[%s] err: %s`, lConfig.TimeZone, err))
	}

	cfg.EncoderConfig = zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "M",
		StacktraceKey: "S",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.In(location).Format(lConfig.TimeLayout))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   FullCallerEncoder, // zapcore.FullCallerEncoder, // 长路径编码器 FullCallerEncoder 短路径编码器 ShortCallerEncoder
	}

	encoder = zapcore.NewConsoleEncoder(cfg.EncoderConfig)

	//	if lConfig.FilePath != `` {
	//		lumberjackLogger := &lumberjack.Logger{
	//			Filename:   lConfig.FilePath + ".log", // 日志文件路径
	//			MaxSize:    lConfig.Rotate.MaxSize,    // megabytes
	//			MaxBackups: lConfig.Rotate.MaxBackups, //
	//			MaxAge:     lConfig.Rotate.MaxAge,     // days
	//			Compress:   true,
	//  }

	// 输出到日志文件 - 考虑到使用场景 基本用不到记录到文件中 - 故暂时注释这些
	// lumberjackLogger := &lumberjack.Logger{
	// 	Filename: lConfig.FilePath + ".log",
	// 	Compress: true,
	// }
	//
	// writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLogger))
	//
	// allCores = append(allCores, zapcore.NewCore(
	// 	encoder,
	// 	writeSyncer,
	// 	cfg.Level,
	// ))

	allCores = append(allCores, zapcore.NewCore(encoder, os.Stdout, cfg.Level))

	allCores = append(allCores, cores...)

	underlyingLogger = zap.New(zapcore.NewTee(allCores...), zap.AddCaller())

	return newLogger(underlyingLogger.With(zap.String(`系统`, `sys`)), ``, 1, true)
}

func FullCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	//fmt.Println("config app:", strings.Replace(caller.String(), global.BasePath, "", 1))
	enc.AppendString(strings.Replace(caller.String(), "/Users/mac/Documents/work_upay/usdtpay/", "", 1))
}

// SkipPanic 是一个 Go 函数，用于从 panic 中恢复并记录错误消息。
func SkipPanic() {
	if s := recover(); s != any(nil) {
		Error("!!!Panic :", zap.Any("err", s))
	}
}
