# goutils

`goutils` 是一个 Go 语言的通用工具库合集，旨在简化日常开发任务。它提供了开箱即用的日志记录、配置管理、信号处理、结构体操作等辅助工具。

## 安装

```bash
go get github.com/reggiepy/goutils/v2
```

## 功能特性与用法

### 1. 日志工具 (`logutil`)

提供两种模式的日志封装，解决了与 `go.uber.org/zap` 包名冲突的问题。

#### 模式一：`logutil/zaputil` (工厂模式)

创建标准的 `*zap.Logger` 并返回清理函数。适用于需要手动管理清理逻辑的场景。

```go
package main

import (
	"github.com/reggiepy/goutils/v2/logutil/zaputil"
	"go.uber.org/zap"
)

func main() {
	// 创建默认日志配置
	config := zaputil.NewLoggerConfig(
		zaputil.WithFile("app.log"),    // 日志文件名
		zaputil.WithLogLevel("debug"),  // 日志级别
		zaputil.WithInConsole(true),    // 是否输出到控制台
		zaputil.WithInFile(true),       // 是否输出到文件
	)

	// 初始化 logger
	logger, cleanup := zaputil.NewLogger(config)
	defer cleanup() // 确保程序退出时刷新缓冲区

	// 使用 logger
	logger.Info("Application started", zap.String("module", "main"))
	logger.Error("An error occurred", zap.Int("code", 500))
}
```

#### 模式二：`logutil/zlog` (封装模式 - 推荐)

提供 `*zlog.Logger` 包装器，内置了生命周期管理（不再返回 cleanup 函数，而是提供 `Close()` 方法）。

```go
package main

import (
	"github.com/reggiepy/goutils/v2/logutil/zlog"
)

func main() {
	// 直接使用选项初始化 logger
	logger := zlog.NewLogger(
		zlog.WithFile("app_v2.log"),
		zlog.WithLogLevel("info"),
		zlog.WithInConsole(true),
	)
	// 安全地使用 defer Close()
	defer logger.Close()

	// 使用 logger (继承了 zap.Logger 的所有方法)
	logger.Info("V2 Logger started")
}
```

### 2. 系统工具 (`sysutil`)

包含优雅停机等系统级操作助手。

```go
package main

import (
	"fmt"
	"time"
	"github.com/reggiepy/goutils/v2/sysutil"
)

func main() {
	// 注册清理任务
	sysutil.OnExit(func() {
		fmt.Println("正在清理数据库连接...")
		time.Sleep(1 * time.Second) // 模拟耗时操作
		fmt.Println("数据库已关闭。")
	})

	sysutil.OnExit(func() {
		fmt.Println("正在停止 HTTP 服务...")
	})

	// 启动你的应用程序逻辑...
	fmt.Println("程序运行中。按 Ctrl+C 退出。")

	// 阻塞并等待退出信号（设置 5秒 超时时间用于清理）
	sysutil.WaitExit(5 * time.Second)
}
```

### 3. 结构体工具 (`structutil`)

用于结构体操作、深度判空检查和 Map 转换的辅助工具。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/structutil"
)

type User struct {
	Name string
	Age  int
}

type UserDTO struct {
	Name string
	Role string
}

func main() {
	user := &User{Name: "Alice", Age: 30}
	dto := &UserDTO{}

	// 拷贝同名且同类型的字段
	structutil.CopyIntersectionStruct(user, dto)
	fmt.Println(dto.Name) // 输出: Alice

	// 检查结构体是否为空（递归深度检查）
	emptyUser := &User{}
	fmt.Println(structutil.IsStructEmpty(emptyUser)) // 输出: true
}
```

### 4. 数组工具 (`arrutil`)

通用的数组/切片操作及 Set 实现。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/arrutil"
)

func main() {
	// 检查切片中是否存在某元素
	nums := []int{1, 2, 3}
	if arrutil.InArray(2, nums) {
		fmt.Println("找到 2 了!")
	}

	// 字符串 Set (集合)
	set := arrutil.NewSet("a", "b")
	set.Add("c")
	if set.Has("a") {
		fmt.Println("集合中包含 'a'")
	}
}
```

### 5. 版本工具 (`verutil`)

语义化版本比较助手。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/verutil"
)

func main() {
	v1 := "1.2.3"
	v2 := "1.3.0"

	// 比较版本号
	if verutil.LessThan(v1, v2) {
		fmt.Printf("%s 早于 %s\n", v1, v2)
	}
}
```

### 6. YAML 工具 (`yamlutil`)

用于读写 YAML 文件的简单封装。

```go
package main

import "github.com/reggiepy/goutils/v2/yamlutil"

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func main() {
	cfg := Config{Host: "localhost", Port: 8080}
	
	// 写入文件
	yamlutil.WriteFile("config.yaml", cfg)

	// 读取文件
	var loadedCfg Config
	yamlutil.ReadFile("config.yaml", &loadedCfg)
}
```

### 7. 枚举工具 (`enumutil`)

基于字符串的枚举验证工具。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/enumutil"
)

func main() {
	// 定义允许的值和默认值
	status := enumutil.NewEnum([]string{"pending", "active", "closed"}, "pending")

	// 设置值（带验证）
	err := status.Set("active")
	if err != nil {
		fmt.Println("错误:", err)
	}
	
	fmt.Println("当前状态:", status.String())
}
```

## 开发

### 运行测试

```bash
go test ./...
```

### 构建

```bash
go build ./...
```