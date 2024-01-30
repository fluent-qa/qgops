package logging

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type ZapLoggerOptions struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Folder        string `mapstructure:"folder" json:"folder"  yaml:"folder"`                        // 日志文件夹
	FileName      string `mapstructure:"file-name" json:"file-name"  yaml:"file-name"`               // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名

	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	MaxSize      int  `mapstructure:"max-size" json:"max-size" yaml:"max-size"`                   // 日志大小
	MaxBackups   int  `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`          // 日志大小
}

var DefaultZapLoggingOptions = ZapLoggerOptions{
	MaxSize:       500, // megabytes
	MaxBackups:    3,
	MaxAge:        28, // days
	FileName:      "app",
	Level:         "info",
	Format:        "console",
	Prefix:        "[fluent-go]",
	Folder:        "logging",
	ShowLine:      true,
	EncodeLevel:   "LowercaseColorLevelEncoder",
	StacktraceKey: "stacktrace",
	LogInConsole:  true,
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *ZapLoggerOptions) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *ZapLoggerOptions) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
