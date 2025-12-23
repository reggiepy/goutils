package zapLogger

type Option interface {
	apply(config *LoggerConfig)
}

type optionFunc func(config *LoggerConfig)

func (o optionFunc) apply(config *LoggerConfig) {
	o(config)
}

// WithFile 设置日志文件名
func WithFile(file string) Option {
	return optionFunc(func(c *LoggerConfig) { c.File = file })
}

// WithMaxSize 设置日志文件大小限制，单位为 MB
func WithMaxSize(maxSize int) Option {
	return optionFunc(func(c *LoggerConfig) { c.MaxSize = maxSize })
}

// WithMaxBackups 设置最大保留的旧日志文件数量
func WithMaxBackups(maxBackups int) Option {
	return optionFunc(func(c *LoggerConfig) { c.MaxBackups = maxBackups })
}

// WithMaxAge 设置旧日志文件保留天数
func WithMaxAge(maxAge int) Option {
	return optionFunc(func(c *LoggerConfig) { c.MaxAge = maxAge })
}

// WithCompress 设置是否压缩旧日志文件
func WithCompress(compress bool) Option {
	return optionFunc(func(c *LoggerConfig) { c.Compress = compress })
}

// WithLogLevel 设置日志等级
func WithLogLevel(level string) Option {
	return optionFunc(func(c *LoggerConfig) { c.Level = level })
}

// WithLogFormat 设置日志格式
func WithLogFormat(format string) Option {
	return optionFunc(func(c *LoggerConfig) { c.Format = format })
}

// WithInConsole 设置是否输出到终端
func WithInConsole(InConsole bool) Option {
	return optionFunc(func(c *LoggerConfig) { c.InConsole = InConsole })
}

// WithInFile 设置是否输出到文件
func WithInFile(InFile bool) Option {
	return optionFunc(func(c *LoggerConfig) { c.InFile = InFile })
}

// WithReplaceGlobals 设置是否替换全局日志记录器
func WithReplaceGlobals(replaceGlobals bool) Option {
	return optionFunc(func(c *LoggerConfig) { c.ReplaceGlobals = replaceGlobals })
}

// WithCaller 设置是否显示调用者信息
func WithCaller(Caller bool) Option {
	return optionFunc(func(c *LoggerConfig) { c.Caller = Caller })
}

// WithCallerSkip 设置调用者信息跳过的层级
func WithCallerSkip(skip int) Option {
	return optionFunc(func(c *LoggerConfig) { c.CallerSkip = skip })
}

// WithStacktraceLevel 设置堆栈跟踪日志级别
func WithStacktraceLevel(level string) Option {
	return optionFunc(func(c *LoggerConfig) { c.StacktraceLevel = level })
}
