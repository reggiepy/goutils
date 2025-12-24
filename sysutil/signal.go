package sysutil

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	shutdownHooks []func()
	mu            sync.Mutex
	once          sync.Once
	wg            sync.WaitGroup

	sigChan    chan os.Signal
	sigOnce    sync.Once
	logHandler = defaultLogHandler
)

// defaultLogHandler 默认日志处理
func defaultLogHandler(msg string) {
	fmt.Printf("[Exit] %s\n", msg)
}

// SetExitMessageHandler 设置退出日志处理函数
func SetExitMessageHandler(handler func(msg string)) {
	if handler != nil {
		logHandler = handler
	}
}

// OnExit 注册退出钩子函数
func OnExit(fnExit func()) {
	mu.Lock()
	defer mu.Unlock()
	shutdownHooks = append(shutdownHooks, fnExit)
}

// TriggerExitSignal 手动触发退出信号（例如业务主动退出）
func TriggerExitSignal() {
	initSigChan()
	go func() {
		sigChan <- syscall.SIGTERM
	}()
}

// ExecuteShutdownHooks 执行所有退出钩子函数，只执行一次
func ExecuteShutdownHooks() {
	logHandler("准备退出，开始执行清理操作")
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock()

		for i := len(shutdownHooks) - 1; i >= 0; i-- {
			wg.Add(1)
			go func(fn func()) {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						logHandler("退出钩子 panic：" + fmt.Sprintf("%v", r))
					}
				}()
				fn()
			}(shutdownHooks[i])
		}
	})
}

// WaitExit 等待退出信号，并在退出时执行钩子，带超时保护
func WaitExit(timeout time.Duration) {
	initSigChan()
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(sigChan)

	logHandler("等待退出信号 (Ctrl+C / SIGTERM)...")
	sig := <-sigChan
	logHandler("收到退出信号：" + sig.String())

	ExecuteShutdownHooks()

	// 等待退出清理
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		logHandler("清理完成，程序正常退出")
	case <-time.After(timeout):
		logHandler("退出处理因超时中止")
	}
}

// --- 私有辅助函数 ---

func initSigChan() {
	sigOnce.Do(func() {
		sigChan = make(chan os.Signal, 1)
	})
}
