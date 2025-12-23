package zapLogger

import "encoding/json"

type LoggerConfig struct {
	InFile         bool `json:"InFile" yaml:"InFile"`                 // 是否输出到文件
	InConsole      bool `json:"InConsole" yaml:"InConsole"`           // 是否输出到终端
	ReplaceGlobals bool `json:"ReplaceGlobals" yaml:"ReplaceGlobals"` // 是否替换全局日志记录器

	File            string `json:"File" yaml:"File"`                       // 日志文件名
	MaxSize         int    `json:"MaxSize" yaml:"MaxSize"`                 // 日志文件大小限制（单位：MB）
	MaxBackups      int    `json:"MaxBackups" yaml:"MaxBackups"`           // 最大保留的旧日志文件数量
	MaxAge          int    `json:"MaxAge" yaml:"MaxAge"`                   // 旧日志文件保留天数
	Compress        bool   `json:"Compress" yaml:"Compress"`               // 是否压缩旧日志文件
	Level           string `json:"LogLevel" yaml:"LogLevel"`               // 日志级别
	Format          string `json:"LogFormat" yaml:"LogFormat"`             // 日志格式（如：json、logfmt）
	Caller          bool   `json:"Caller" yaml:"Caller"`                   // 是否显示调用者信息
	CallerSkip      int    `json:"CallerSkip" yaml:"CallerSkip"`           // 调用者信息跳过的层级
	StacktraceLevel string `json:"StacktraceLevel" yaml:"StacktraceLevel"` // 堆栈跟踪日志级别
}

func (l *LoggerConfig) clone() *LoggerConfig {
	clone := *l
	return &clone
}

func (l *LoggerConfig) ToJSON() string {
	jsonStr, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

func (l *LoggerConfig) LoadJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), l)
}

func (l *LoggerConfig) WithOptions(options ...Option) *LoggerConfig {
	c := l.clone()
	for _, opt := range options {
		opt.apply(c)
	}
	return c
}

// NewLoggerConfig 创建默认配置
func NewLoggerConfig(opts ...Option) *LoggerConfig {
	config := &LoggerConfig{
		InFile:          true,      // 默认输出到文件
		InConsole:       false,     // 默认不输出到控制台
		File:            "app.log", // 默认日志文件名
		MaxSize:         1,         // 默认文件大小限制 1MB
		MaxBackups:      5,         // 默认保留 5 个旧文件
		MaxAge:          30,        // 默认保留 30 天
		Compress:        false,     // 默认不压缩
		Level:           "info",    // 默认日志级别为 info
		Format:          "json",    // 默认使用 JSON 格式
		ReplaceGlobals:  false,     // 默认不替换全局 zap logger
		Caller:          true,      // 默认开启调用者信息
		CallerSkip:      1,         // 默认跳过 1 层调用栈
		StacktraceLevel: "panic",   // 默认 panic 级别才打印堆栈信息
	}
	return config.WithOptions(opts...)
}
