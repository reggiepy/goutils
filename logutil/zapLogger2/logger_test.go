package zapLogger2

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	// Test default
	logger := NewLogger()
	defer logger.Close()

	assert.NotNil(t, logger)
	assert.NotNil(t, logger.Config)
	logger.Info("Default v2 logger test")
}

func TestNewLoggerWithOptions(t *testing.T) {
	tmpFile := "test_v2.log"
	defer os.Remove(tmpFile)

	// v2 Style: Pass options directly to NewLogger
	logger := NewLogger(
		WithFile(tmpFile),
		WithInFile(true),
		WithInConsole(false),
		WithLogLevel("debug"),
		WithCallerSkip(2),
	)
	defer func(logger *Logger) {
		_ = logger.Close()
	}(logger)

	assert.NotNil(t, logger)
	logger.Debug("v2 file test log")

	_, err := os.Stat(tmpFile)
	assert.NoError(t, err)
}

func TestNewLoggerWithZapOptions(t *testing.T) {
	var called bool
	hook := func(entry zapcore.Entry) error {
		called = true
		return nil
	}

	logger := NewLogger(
		WithZapOptions(zap.Hooks(hook)),
	)
	defer logger.Close()

	logger.Info("Testing hook v2")
	logger.Sync()

	assert.True(t, called)
}

func TestNewLoggerWithConfigStruct(t *testing.T) {
	// Simulate loading from JSON
	jsonConfig := `{
		"File": "config_test.log",
		"LogLevel": "warn",
		"InConsole": true
	}`

	var cfg Config
	err := json.Unmarshal([]byte(jsonConfig), &cfg)
	assert.NoError(t, err)

	// Use WithConfig to pass the struct
	logger := NewLogger(
		WithConfig(cfg),
		WithMaxSize(10), // Override or add to the config
	)
	defer logger.Close()

	assert.NotNil(t, logger)
	logger.Warn("Config struct test")
}
