package loghelp

import "go.uber.org/zap/zapcore"

// LogConfig 日志器配置
type LogConfig struct {
	File       LogFile   `yaml:"File" json:"file"`              // 日志文件
	Switch     LogSwitch `yaml:"Switch" json:"switch"`          // 日志开关
	TimeZone   string    `yaml:"TimeZone" json:"time_zone"`     // 时区，默认defaultTimeZone,可以从 https://www.zeitverschiebung.net/en/ 查询时区信息
	TimeLayout string    `yaml:"TimeLayout" json:"time_layout"` // 输出时间格式,默认为defaultTimeLayout,任何Go支持的格式都是合法的
	LogLevel   LogLevel  `yaml:"LogLevel" json:"log_level"`     // 日志级别
}

// LogLevel 日志级别
type LogLevel struct {
	Default  zapcore.Level            `yaml:"Default" json:"default"`   // 默认
	Specific map[string]zapcore.Level `yaml:"Specific" json:"specific"` // 具体
}

// LogFile 日志文件级别
type LogFile struct {
	Info    string `json:"file_info" yaml:"Info"`       // 日志文件(info)
	Error   string `json:"file_error" yaml:"Error"`     // 日志文件(error)
	Mysql   string `json:"file_mysql" yaml:"Mysql"`     // 日志文件(mysql)
	Service string `json:"file_service" yaml:"Service"` // 日志文件(service)
	ReqRet  string `json:"file_req_ret" yaml:"ReqRet"`  // 日志文件(req_ret)
}

// LogSwitch 日志打印开关
type LogSwitch struct {
	InfoWrite  bool `json:"switch_info_write" yaml:"InfoWrite"`   // 是否写文件(info)
	InfoPrint  bool `json:"switch_info_print" yaml:"InfoPrint"`   // 是否打印显示出来(info)
	DebugWrite bool `json:"switch_debug_write" yaml:"DebugWrite"` // 是否写文件(debug)
	DebugPrint bool `json:"switch_debug_print" yaml:"DebugPrint"` // 是否打印显示出来(debug)
	ErrorWrite bool `json:"switch_error_write" yaml:"ErrorWrite"` // 是否写文件(error)
	ErrorPrint bool `json:"switch_error_print" yaml:"ErrorPrint"` // 是否打印显示出来(error)
}
