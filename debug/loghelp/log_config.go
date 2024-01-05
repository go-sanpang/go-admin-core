package loghelp

import "go.uber.org/zap/zapcore"

// LogConfig 日志器配置
type LogConfig struct {
	TimeZone   string   `yaml:"TimeZone" json:"time_zone"`     // 时区，默认defaultTimeZone,可以从 https://www.zeitverschiebung.net/en/ 查询时区信息
	TimeLayout string   `yaml:"TimeLayout" json:"time_layout"` // 输出时间格式,默认为defaultTimeLayout,任何Go支持的格式都是合法的
	LogLevel   LogLevel `yaml:"LogLevel" json:"log_level"`     // 日志级别
}

// LogLevel 日志级别
type LogLevel struct {
	Default  zapcore.Level            `yaml:"Default" json:"default"`   // 默认
	Specific map[string]zapcore.Level `yaml:"Specific" json:"specific"` // 具体
}
