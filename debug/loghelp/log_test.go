package loghelp

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestName(t *testing.T) {
	Error(`主程序`)

	// l.Derive(`main`).Error(`<<<<<<<<启动程序>>>>>>>>`)
	l := GetLogger(`主程序`, `main`)

	p := l.SetLevel(zapcore.DebugLevel)
	p.Debug(`<<<<<<<<启动程序>>>>>>>>`)
	p.Info(`<<<<<<<<启动程序>>>>>>>>`)
	p.Error(`<<<<<<<<启动程序>>>>>>>>`)
	return
	GetLogger(`主程序`, `main`).Error(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`).Error(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`).Panic(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`).Fatal(`<<<<<<<<启动程序>>>>>>>>`)
	GetLogger(`主程序?`, `mai2n`).Error(`<<<<<<<<启动程序>>>>>>>>`)
}
