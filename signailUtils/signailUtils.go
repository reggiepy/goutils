package signailUtils

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	shutdownHooks      []func()         // 退出时需要执行的钩子函数列表
	exitMessageHandler func(msg string) // 退出消息处理函数
	mu                 sync.Mutex       // 确保并发情况下对 shutdownHooks 的操作安全
	once               sync.Once        // 确保退出函数只执行一次
	wg                 sync.WaitGroup   // 确保所有协程都在程序退出前完成
	shutdownStarted    bool
)

// SetExitMessageHandler - 设置自定义的退出消息处理函数
func SetExitMessageHandler(handler func(msg string)) {
	exitMessageHandler = handler // 将自定义处理函数赋值
}

// GetExitMessageHandler - 获取退出消息处理函数，默认为标准输出
func GetExitMessageHandler() func(msg string) {
	if exitMessageHandler != nil {
		return exitMessageHandler // 返回自定义的处理函数
	}
	// 默认处理函数，打印消息到控制台
	return func(msg string) {
		fmt.Println(msg)
	}
}

// OnExit - 注册退出处理函数
func OnExit(fnExit func()) {
	mu.Lock()
	defer mu.Unlock()
	if shutdownStarted {
		panic("shutdown already in progress, can't register new hook")
	}
	shutdownHooks = append(shutdownHooks, fnExit)
	wg.Add(1)
}

// ExecuteShutdownHooks - 执行所有注册的退出钩子函数，只执行一次
func ExecuteShutdownHooks() {
	GetExitMessageHandler()("准备退出,执行清理操作")
	once.Do(func() { // 确保只执行一次
		mu.Lock()                                      // 获取锁以确保安全
		defer mu.Unlock()                              // 解锁
		for i := len(shutdownHooks) - 1; i >= 0; i-- { // 倒序执行
			go func(fn func()) {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						GetExitMessageHandler()(logPanic(r))
					}
				}()
				fn()
			}(shutdownHooks[i])
		}
	})
	GetExitMessageHandler()("程序清理完毕。")
}

// WaitExit - 同步等待退出信号后退出，可指定自定义退出处理函数
func WaitExit(timeout time.Duration) {
	sigChan := make(chan os.Signal, 1)                      // 创建信号通道
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT) // 监听 SIGTERM 和 SIGINT 信号

	GetExitMessageHandler()(logSignal(<-sigChan))// 阻塞等待信号
	ExecuteShutdownHooks() // 执行已注册的退出钩子函数

	// 等待所有退出钩子完成或超时
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done) // 退出钩子执行完，关闭通道
	}()

	select {
	case <-done: // 等待完成
		GetExitMessageHandler()("所有清理已完成，程序正常退出。")
	case <-time.After(timeout): // 设置超时时间
		GetExitMessageHandler()("退出处理因超时而中止。")
	}
}

// 工具函数：格式化 panic 日志
func logPanic(r interface{}) string {
	return "退出钩子发生 panic: " + formatAny(r)
}

// 工具函数：格式化信号日志
func logSignal(sig os.Signal) string {
	return "收到系统信号: " + sig.String()
}

// 工具函数：格式化 interface{}
func formatAny(v interface{}) string {
	return fmt.Sprintf("%v", v)
}