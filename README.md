# goutils

[中文说明](./README_ZH.md)

`goutils` is a collection of common utility packages for Go, designed to simplify daily development tasks. It provides ready-to-use helpers for logging, configuration management, signal handling, struct manipulation, and more.

## Installation

```bash
go get github.com/reggiepy/goutils/v2
```

## Features & Usage

### 1. Logging

#### V1: `logutil/zaputil` (Factory Pattern)

Creates a standard `*zap.Logger` and returns a cleanup function.

```go
package main

import (
	"github.com/reggiepy/goutils/v2/logutil/zaputil"
	"go.uber.org/zap"
)

func main() {
	// Create a default logger config
	config := zaputil.NewLoggerConfig(
		zaputil.WithFile("app.log"),
		zaputil.WithLogLevel("debug"),
		zaputil.WithInConsole(true),
		zaputil.WithInFile(true),
	)

	// Initialize logger
	logger, cleanup := zaputil.NewLogger(config)
	defer cleanup()

	// Use the logger
	logger.Info("Application started", zap.String("module", "main"))
	logger.Error("An error occurred", zap.Int("code", 500))
}
```

#### V2: `logutil/zlog` (Wrapper Pattern)

Provides a `*zlog.Logger` wrapper that manages its own lifecycle (no cleanup function returned).

```go
package main

import (
	"github.com/reggiepy/goutils/v2/logutil/zlog"
)

func main() {
	// Initialize logger directly with options
	logger := zlog.NewLogger(
		zlog.WithFile("app_v2.log"),
		zlog.WithLogLevel("info"),
		zlog.WithInConsole(true),
	)
	// Safe to defer Close()
	defer logger.Close()

	// Use the logger (inherits methods from zap.Logger)
	logger.Info("V2 Logger started")
}
```

#### Interface: `logutil/logx` & Adapter: `logutil/adapters/zapx`

A minimal, dependency-agnostic logger interface `logx.Logger` and its Zap adapter `zapx.ZapLogger`. This allows your domain logic to depend on `logx.Logger` interface instead of concrete Zap implementation.

```go
package main

import (
	"github.com/reggiepy/goutils/v2/logutil/logx"
	"github.com/reggiepy/goutils/v2/logutil/adapters/zapx"
	"go.uber.org/zap"
)

// Domain logic depends on minimal interface
func RunBusinessLogic(logger logx.Logger) {
	logger.Info("Business logic started", "module", "core")
	// ...
}

func main() {
	// Setup Zap
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	// Adapt to logx.Logger
	logger := zapx.NewZapLogger(zapLogger)

	RunBusinessLogic(logger)
}
```

### 2. System Utilities (`sysutil`)

Helpers for system-level operations, including graceful shutdown handling.

```go
package main

import (
	"fmt"
	"time"
	"github.com/reggiepy/goutils/v2/sysutil"
)

func main() {
	// Register cleanup tasks
	sysutil.OnExit(func() {
		fmt.Println("Cleaning up database connections...")
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Println("Database closed.")
	})

	sysutil.OnExit(func() {
		fmt.Println("Stopping HTTP server...")
	})

	// Start your application logic...
	fmt.Println("App is running. Press Ctrl+C to exit.")

	// Block and wait for exit signal (with 5s timeout for cleanup)
	sysutil.WaitExit(5 * time.Second)
}
```

### 3. Struct Utilities (`structutil`)

Helpers for struct manipulation, deep empty checks, and map conversion.

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

	// Copy fields with same name and type
	structutil.CopyIntersectionStruct(user, dto)
	fmt.Println(dto.Name) // Output: Alice

	// Check if struct is empty (deep check)
	emptyUser := &User{}
	fmt.Println(structutil.IsStructEmpty(emptyUser)) // Output: true
}
```

### 4. Array Utilities (`arrutil`)

Common array/slice operations and a Set implementation.

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/arrutil"
)

func main() {
	// Check if item exists in slice
	nums := []int{1, 2, 3}
	if arrutil.InArray(2, nums) {
		fmt.Println("Found 2!")
	}

	// String Set
	set := arrutil.NewSet("a", "b")
	set.Add("c")
	if set.Has("a") {
		fmt.Println("Set has 'a'")
	}
}
```

### 5. Version Utilities (`verutil`)

Semantic versioning comparison helpers.

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/verutil"
)

func main() {
	v1 := "1.2.3"
	v2 := "1.3.0"

	if verutil.LessThan(v1, v2) {
		fmt.Printf("%s is older than %s\n", v1, v2)
	}
}
```

### 6. YAML Utilities (`yamlutil`)

Simple wrapper for reading and writing YAML files.

```go
package main

import "github.com/reggiepy/goutils/v2/yamlutil"

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func main() {
	cfg := Config{Host: "localhost", Port: 8080}
	
	// Write
	yamlutil.WriteFile("config.yaml", cfg)

	// Read
	var loadedCfg Config
	yamlutil.ReadFile("config.yaml", &loadedCfg)
}
```

### 7. Enum Utilities (`enumutil`)

String-based enum validation.

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/v2/enumutil"
)

func main() {
	// Define allowed values
	status := enumutil.NewEnum([]string{"pending", "active", "closed"}, "pending")

	// Set value with validation
	err := status.Set("active")
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	fmt.Println("Current Status:", status.String())
}
```

## Development

### Run Tests

```bash
go test ./...
```

### Build

```bash
go build ./...
```
