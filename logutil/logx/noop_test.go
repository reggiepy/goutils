package logx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoopLogger(t *testing.T) {
	as := assert.New(t)
	l := NewNoopLogger()

	// These should not panic or do anything
	l.Debug("test", "key", "val")
	l.Info("test")
	l.Warn("test")
	l.Error("test")

	// With should return the same logger (or another noop logger)
	child := l.With("foo", "bar")
	as.NotNil(child)
	child.Info("child test")

	// Test Discard
	as.NotNil(Discard)
	Discard.Info("discard test")
}
