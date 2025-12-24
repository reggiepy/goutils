package zaputil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	// Test default config
	config := NewLoggerConfig()
	logger, cleanup := NewLogger(config)
	defer cleanup()

	assert.NotNil(t, logger)
	logger.Info("Default config test")
}

func TestNewLoggerWithFile(t *testing.T) {
	// Temporary file for testing
	tmpFile := "test_log.log"
	defer os.Remove(tmpFile)

	config := NewLoggerConfig(
		WithFile(tmpFile),
		WithInFile(true),
		WithInConsole(false),
		WithLogLevel("debug"),
	)

	logger, cleanup := NewLogger(config)
	defer cleanup()

	assert.NotNil(t, logger)
	logger.Debug("File test log")

	// Verify file was created
	_, err := os.Stat(tmpFile)
	assert.NoError(t, err)
}

func TestNewLoggerWithExtraOptions(t *testing.T) {
	config := NewLoggerConfig()

	// Use a hook to verify the extra option works
	var called bool
	hook := func(entry zapcore.Entry) error {
		called = true
		return nil
	}

	logger, cleanup := NewLogger(config, zap.Hooks(hook))
	defer cleanup()

	logger.Info("Testing hook")

	// Force sync to ensure hook runs (though hooks usually run synchronously)
	logger.Sync()

	assert.True(t, called, "Hook should have been called")
}

func TestLoggerConfig_ToJSON_LoadJSON(t *testing.T) {
	config := NewLoggerConfig()
	config.Level = "error"

	jsonStr := config.ToJSON()
	assert.Contains(t, jsonStr, `"LogLevel":"error"`)

	newConfig := &LoggerConfig{}
	err := newConfig.LoadJSON(jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, "error", newConfig.Level)
}

func TestNewLoggerWithAdvancedOptions(t *testing.T) {
	config := NewLoggerConfig(
		WithCallerSkip(1),
		WithStacktraceLevel("error"),
	)

	logger, cleanup := NewLogger(config)
	defer cleanup()

	assert.NotNil(t, logger)
	assert.Equal(t, 1, config.CallerSkip)
	assert.Equal(t, "error", config.StacktraceLevel)
}
