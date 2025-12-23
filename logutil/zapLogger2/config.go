package zapLogger2

import (
	"encoding/json"

	"go.uber.org/zap"
)

type Config struct {
	InFile          bool   `json:"InFile" yaml:"InFile"`                   // 是否输出到文件
	InConsole       bool   `json:"InConsole" yaml:"InConsole"`             // 是否输出到终端
	ReplaceGlobals  bool   `json:"ReplaceGlobals" yaml:"ReplaceGlobals"`   // 是否替换全局日志记录器
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

	ZapOptions []zap.Option `json:"-" yaml:"-"` // 额外的 zap 选项
}

// Option 定义配置选项函数
type Option func(config *Config)

// ToJSON 序列化配置
func (l *Config) ToJSON() string {
	jsonStr, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

// LoadJSON 反序列化配置
func (l *Config) LoadJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), l)
}

// NewConfig 创建并返回默认配置对象
func NewConfig() *Config {
	return &Config{
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
		ZapOptions:      make([]zap.Option, 0),
	}
}

// --- Options ---

// WithFile 设置日志文件名
func WithFile(file string) Option {
	return func(c *Config) { c.File = file }
}

// WithMaxSize 设置日志文件大小限制，单位为 MB
func WithMaxSize(maxSize int) Option {
	return func(c *Config) { c.MaxSize = maxSize }
}

// WithMaxBackups 设置最大保留的旧日志文件数量
func WithMaxBackups(maxBackups int) Option {
	return func(c *Config) { c.MaxBackups = maxBackups }
}

// WithMaxAge 设置旧日志文件保留天数
func WithMaxAge(maxAge int) Option {
	return func(c *Config) { c.MaxAge = maxAge }
}

// WithCompress 设置是否压缩旧日志文件
func WithCompress(compress bool) Option {
	return func(c *Config) { c.Compress = compress }
}

// WithLogLevel 设置日志等级
func WithLogLevel(level string) Option {
	return func(c *Config) { c.Level = level }
}

// WithLogFormat 设置日志格式
func WithLogFormat(format string) Option {
	return func(c *Config) { c.Format = format }
}

// WithInConsole 设置是否输出到终端
func WithInConsole(InConsole bool) Option {
	return func(c *Config) { c.InConsole = InConsole }
}

// WithInFile 设置是否输出到文件
func WithInFile(InFile bool) Option {
	return func(c *Config) { c.InFile = InFile }
}

// WithReplaceGlobals 设置是否替换全局日志记录器
func WithReplaceGlobals(replaceGlobals bool) Option {
	return func(c *Config) { c.ReplaceGlobals = replaceGlobals }
}

// WithCaller 设置是否显示调用者信息
func WithCaller(Caller bool) Option {
	return func(c *Config) { c.Caller = Caller }
}

// WithCallerSkip 设置调用者信息跳过的层级
func WithCallerSkip(skip int) Option {
	return func(c *Config) { c.CallerSkip = skip }
}

// WithStacktraceLevel 设置堆栈跟踪日志级别
func WithStacktraceLevel(level string) Option {
	return func(c *Config) { c.StacktraceLevel = level }
}

// WithZapOptions 添加额外的 zap 选项
func WithZapOptions(opts ...zap.Option) Option {
	return func(c *Config) {
		c.ZapOptions = append(c.ZapOptions, opts...)
	}
}

// WithConfig 使用已有的配置对象覆盖当前配置
func WithConfig(cfg Config) Option {
	return func(c *Config) {
		// 保存原有的 ZapOptions，因为 cfg 中可能没有（如果是从 json 加载的）
		currentZapOptions := c.ZapOptions
		*c = cfg

		// 合并 ZapOptions
		if len(cfg.ZapOptions) > 0 {
			c.ZapOptions = cfg.ZapOptions
		} else {
			c.ZapOptions = currentZapOptions
		}
	}
}
