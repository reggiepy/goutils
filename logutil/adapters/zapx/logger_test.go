package zapx

import (
	"testing"

	"github.com/reggiepy/goutils/v2/logutil/logx"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	as := assert.New(t)

	// Create an observer core to verify logs
	core, logs := observer.New(zapcore.DebugLevel)
	logger := zap.New(core)

	// Create the adapter
	var l logx.Logger = NewZapLogger(logger)

	// Test Debug
	l.Debug("debug msg", "key", "val")
	as.Equal(1, logs.Len())
	as.Equal("debug msg", logs.All()[0].Message)
	as.Equal(zapcore.DebugLevel, logs.All()[0].Level)
	as.Equal("val", logs.All()[0].ContextMap()["key"])

	// Test Info
	l.Info("info msg")
	as.Equal(2, logs.Len())
	as.Equal("info msg", logs.All()[1].Message)
	as.Equal(zapcore.InfoLevel, logs.All()[1].Level)

	// Test With (Context)
	child := l.With("childKey", "childVal")
	child.Warn("child warn")
	as.Equal(3, logs.Len())
	as.Equal("child warn", logs.All()[2].Message)
	as.Equal(zapcore.WarnLevel, logs.All()[2].Level)
	as.Equal("childVal", logs.All()[2].ContextMap()["childKey"])

	// Ensure parent logger is not affected
	l.Error("parent error")
	as.Equal(4, logs.Len())
	as.Nil(logs.All()[3].ContextMap()["childKey"])
}
