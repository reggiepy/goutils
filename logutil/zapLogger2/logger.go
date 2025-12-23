package zapLogger2

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 封装了 zap.Logger 和其底层的配置/资源
type Logger struct {
	*zap.Logger                    // 底层 zap 日志实例
	Config      *Config            // 日志配置信息
	lj          *lumberjack.Logger // 日志轮转器
	cores       []zapcore.Core     // 日志核心列表
}

// Close 安全关闭日志记录器，刷新缓冲并关闭文件句柄
func (l *Logger) Close() error {
	// 尝试刷新缓冲区
	_ = l.Logger.Sync()

	// 关闭 lumberjack logger（如果存在）
	if l.lj != nil {
		return l.lj.Close()
	}
	return nil
}

// NewLogger 初始化Logger
// v2 改进：返回 *Logger 结构体，包含 Close 方法，不再返回 cleanup 函数
func NewLogger(opts ...Option) *Logger {
	// 1. 初始化默认配置
	config := NewConfig()

	// 2. 应用选项
	for _, opt := range opts {
		opt(config)
	}

	// 3. 解析基础配置
	logFormat := getLogFormat(config)
	level := getLogLevel(config)

	var cores []zapcore.Core

	// 4. 构建文件输出 Core
	fileCore, lj := buildFileCore(config, logFormat, level)
	if fileCore != nil {
		cores = append(cores, fileCore)
	}

	// 5. 构建终端输出 Core
	consoleCore := buildConsoleCore(config, logFormat, level)
	if consoleCore != nil {
		cores = append(cores, consoleCore)
	}

	// 6. 创建最终的 Logger
	core := zapcore.NewTee(cores...)
	options := buildZapOptions(config)
	zapLogger := zap.New(core, options...)

	if config.ReplaceGlobals {
		zap.ReplaceGlobals(zapLogger)
	}

	return &Logger{
		Logger: zapLogger,
		Config: config,
		lj:     lj,
		cores:  cores,
	}
}

// buildFileCore 构建文件日志输出核心
func buildFileCore(config *Config, logFormat string, level zapcore.Level) (zapcore.Core, *lumberjack.Logger) {
	if !config.InFile || config.File == "" {
		return nil, nil
	}

	lj := &lumberjack.Logger{
		Filename:   config.File,
		MaxSize:    config.MaxSize,    // 文件大小 (MB)
		MaxBackups: config.MaxBackups, // 保留的最大文件个数
		MaxAge:     config.MaxAge,     // 保留的最大天数
		Compress:   config.Compress,   // 是否压缩
	}

	core := zapcore.NewCore(
		NewEncoder(logFormat),
		zapcore.AddSync(lj), // 将日志写入文件
		level,
	)

	return core, lj
}

// buildConsoleCore 构建终端日志输出核心
func buildConsoleCore(config *Config, logFormat string, level zapcore.Level) zapcore.Core {
	if !config.InConsole {
		return nil
	}

	return zapcore.NewCore(
		NewEncoder(logFormat),
		zapcore.AddSync(os.Stdout), // 输出到终端
		level,
	)
}

// buildZapOptions 构建 zap 选项
func buildZapOptions(config *Config) []zap.Option {
	var options []zap.Option

	if config.Caller {
		options = append(options, zap.AddCaller())
	}
	if config.CallerSkip > 0 {
		options = append(options, zap.AddCallerSkip(config.CallerSkip))
	}

	// 设置堆栈跟踪级别
	stackLevel, err := zapcore.ParseLevel(config.StacktraceLevel)
	if err == nil {
		options = append(options, zap.AddStacktrace(stackLevel))
	}

	// 添加额外的 zap 选项
	if len(config.ZapOptions) > 0 {
		options = append(options, config.ZapOptions...)
	}

	return options
}

func getLogFormat(config *Config) string {
	supportedLogFormats := map[string]struct{}{
		"json":   {},
		"logfmt": {},
	}
	if _, exists := supportedLogFormats[config.Format]; exists {
		return config.Format
	}
	return "json"
}

func getLogLevel(config *Config) zapcore.Level {
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		return zapcore.InfoLevel
	}
	return level
}

func NewEncoder(logFormat string) zapcore.Encoder {
	// 创建一个自定义的 EncoderConfig 实例
	encoderConfig := zap.NewProductionEncoderConfig()

	// 设置时间编码格式为 ISO 8601
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 设置日志级别编码格式为大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 设置时间字段的键名为 "time"
	encoderConfig.TimeKey = "time"

	// 设置持续时间的编码格式为秒
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	// 设置调用者信息的编码格式为简短格式
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	if logFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}
