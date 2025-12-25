package sysutil

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSignalHandling(t *testing.T) {
	as := assert.New(t)

	// Capture logs
	var logs []string
	var logMu sync.Mutex
	SetExitMessageHandler(func(msg string) {
		logMu.Lock()
		defer logMu.Unlock()
		logs = append(logs, msg)
	})

	// Register hook
	hookCalled := false
	var hookMu sync.Mutex
	OnExit(func() {
		hookMu.Lock()
		defer hookMu.Unlock()
		hookCalled = true
	})

	// Trigger signal
	TriggerExitSignal()

	// Wait for exit (should happen quickly due to TriggerExitSignal)
	done := make(chan struct{})
	go func() {
		WaitExit(1 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		// Success
	case <-time.After(2 * time.Second):
		t.Fatal("WaitExit timed out")
	}

	// Verify hook execution
	hookMu.Lock()
	as.True(hookCalled, "Shutdown hook should have been called")
	hookMu.Unlock()

	// Verify logs
	logMu.Lock()
	as.NotEmpty(logs)
	as.Contains(logs[len(logs)-1], "清理完成", "Should log completion")
	logMu.Unlock()
}
