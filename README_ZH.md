# goutils

`goutils` 是一个 Go 语言的通用工具库合集，旨在简化日常开发任务。它提供了开箱即用的日志记录、配置管理、信号处理、结构体操作等辅助工具。

## 安装

```bash
go get github.com/reggiepy/goutils
```

## 功能特性与用法

### 1. 日志工具 (`logutil/zapLogger`)

基于 [uber-go/zap](https://github.com/uber-go/zap) 和 [lumberjack](https://github.com/natefinch/lumberjack) 封装的高性能日志库，支持日志文件自动轮转。

```go
package main

import (
	"github.com/reggiepy/goutils/logutil/zapLogger"
	"go.uber.org/zap"
)

func main() {
	// 创建默认日志配置
	config := zapLogger.NewLoggerConfig(
		zapLogger.WithFile("app.log"),    // 日志文件名
		zapLogger.WithLogLevel("debug"),  // 日志级别
		zapLogger.WithInConsole(true),    // 是否输出到控制台
		zapLogger.WithInFile(true),       // 是否输出到文件
	)

	// 初始化 logger
	logger, cleanup := zapLogger.NewLogger(config)
	defer cleanup() // 确保程序退出时刷新缓冲区

	// 使用 logger
	logger.Info("Application started", zap.String("module", "main"))
	logger.Error("An error occurred", zap.Int("code", 500))
}
```

### 2. 信号处理 (`signailUtils`)

优雅停机助手，用于处理 `SIGINT` 和 `SIGTERM` 信号。

```go
package main

import (
	"fmt"
	"time"
	"github.com/reggiepy/goutils/signailUtils"
)

func main() {
	// 注册清理任务
	signailUtils.OnExit(func() {
		fmt.Println("正在清理数据库连接...")
		time.Sleep(1 * time.Second) // 模拟耗时操作
		fmt.Println("数据库已关闭。")
	})

	signailUtils.OnExit(func() {
		fmt.Println("正在停止 HTTP 服务...")
	})

	// 启动你的应用程序逻辑...
	fmt.Println("程序运行中。按 Ctrl+C 退出。")

	// 阻塞并等待退出信号（设置 5秒 超时时间用于清理）
	signailUtils.WaitExit(5 * time.Second)
}
```

### 3. 结构体工具 (`structUtils`)

用于结构体操作、深度判空检查和 Map 转换的辅助工具。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/structUtils"
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
	structUtils.CopyIntersectionStruct(user, dto)
	fmt.Println(dto.Name) // 输出: Alice

	// 检查结构体是否为空（递归深度检查）
	emptyUser := &User{}
	fmt.Println(structUtils.IsStructEmpty(emptyUser)) // 输出: true
}
```

### 4. 数组工具 (`arrayUtils`)

通用的数组/切片操作及 Set 实现。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/arrayUtils"
)

func main() {
	// 检查切片中是否存在某元素
	nums := []int{1, 2, 3}
	if arrayUtils.InArray(2, nums) {
		fmt.Println("找到 2 了!")
	}

	// 字符串 Set (集合)
	set := arrayUtils.NewSet("a", "b")
	set.Add("c")
	if set.Has("a") {
		fmt.Println("集合中包含 'a'")
	}
}
```

### 5. 版本工具 (`versionUtils`)

语义化版本比较助手。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/versionUtils"
)

func main() {
	v1 := "1.2.3"
	v2 := "1.3.0"

	// 比较版本号
	if versionUtils.LessThan(v1, v2) {
		fmt.Printf("%s 早于 %s\n", v1, v2)
	}
}
```

### 6. YAML 工具 (`yamlutil`)

用于读写 YAML 文件的简单封装。

```go
package main

import "github.com/reggiepy/goutils/yamlutil"

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

### 7. 枚举工具 (`enumUtils`)

基于字符串的枚举验证工具。

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/enumUtils"
)

func main() {
	// 定义允许的值和默认值
	status := enumUtils.NewEnum([]string{"pending", "active", "closed"}, "pending")

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
