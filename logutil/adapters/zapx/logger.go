package zapx

import (
	"github.com/reggiepy/goutils/v2/logutil/logx"
	"go.uber.org/zap"
)

// ZapLogger adapts a zap.SugaredLogger to the logx.Logger interface.
type ZapLogger struct {
	logger *zap.SugaredLogger
}

// NewZapLogger creates a new ZapLogger from a *zap.Logger.
func NewZapLogger(l *zap.Logger) logx.Logger {
	return &ZapLogger{
		logger: l.Sugar(),
	}
}

// NewZapLoggerFromSugared creates a new ZapLogger from a *zap.SugaredLogger.
func NewZapLoggerFromSugared(l *zap.SugaredLogger) logx.Logger {
	return &ZapLogger{
		logger: l,
	}
}

func (l *ZapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *ZapLogger) With(keysAndValues ...interface{}) logx.Logger {
	return &ZapLogger{
		logger: l.logger.With(keysAndValues...),
	}
}
