package loghelp

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestName(t *testing.T) {
	Error(`主程序`)

	cfg := &LogConfig{}

	// l.Derive(`main`).Error(`<<<<<<<<启动程序>>>>>>>>`)
	l := GetLogger(`主程序`, `main`, cfg)

	p := l.SetLevel(zapcore.DebugLevel)
	p.Debug(`<<<<<<<<启动程序>>>>>>>>`)
	p.Info(`<<<<<<<<启动程序>>>>>>>>`)
	p.Error(`<<<<<<<<启动程序>>>>>>>>`)
	// return
	GetLogger(`主程序`, `main`, cfg).Error(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`, cfg).Error(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`, cfg).Panic(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`, cfg).Fatal(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`, cfg).Error(`<<<<<<<<启动程序>>>>>>>>`)
}
