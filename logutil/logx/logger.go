package logx

// Logger defines a minimal logging interface.
// It is designed to be adapter-agnostic, allowing different logging backends (like Zap, Logrus, etc.)
// to be plugged in without changing the application code.
type Logger interface {
	// Debug logs a message at Debug level.
	Debug(msg string, keysAndValues ...interface{})

	// Info logs a message at Info level.
	Info(msg string, keysAndValues ...interface{})

	// Warn logs a message at Warn level.
	Warn(msg string, keysAndValues ...interface{})

	// Error logs a message at Error level.
	Error(msg string, keysAndValues ...interface{})

	// Fatal logs a message at Fatal level and then the process will exit with status set to 1.
	Fatal(msg string, keysAndValues ...interface{})

	// With returns a new Logger instance with the specified key-value pairs attached.
	With(keysAndValues ...interface{}) Logger
}
