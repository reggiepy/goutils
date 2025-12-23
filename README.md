# goutils

[中文说明](./README_ZH.md)

`goutils` is a collection of common utility packages for Go, designed to simplify daily development tasks. It provides ready-to-use helpers for logging, configuration management, signal handling, struct manipulation, and more.

## Installation

```bash
go get github.com/reggiepy/goutils
```

## Features & Usage

### 1. Logging (`logutil/zapLogger`)

A wrapper around [uber-go/zap](https://github.com/uber-go/zap) and [lumberjack](https://github.com/natefinch/lumberjack) for high-performance logging with file rotation support.

```go
package main

import (
	"github.com/reggiepy/goutils/logutil/zapLogger"
	"go.uber.org/zap"
)

func main() {
	// Create a default logger config
	config := zapLogger.NewLoggerConfig(
		zapLogger.WithFile("app.log"),
		zapLogger.WithLogLevel("debug"),
		zapLogger.WithInConsole(true),
		zapLogger.WithInFile(true),
	)

	// Initialize logger
	logger, cleanup := zapLogger.NewLogger(config)
	defer cleanup()

	// Use the logger
	logger.Info("Application started", zap.String("module", "main"))
	logger.Error("An error occurred", zap.Int("code", 500))
}
```

### 2. Signal Handling (`signailUtils`)

Graceful shutdown helper handling `SIGINT` and `SIGTERM`.

```go
package main

import (
	"fmt"
	"time"
	"github.com/reggiepy/goutils/signailUtils"
)

func main() {
	// Register cleanup tasks
	signailUtils.OnExit(func() {
		fmt.Println("Cleaning up database connections...")
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Println("Database closed.")
	})

	signailUtils.OnExit(func() {
		fmt.Println("Stopping HTTP server...")
	})

	// Start your application logic...
	fmt.Println("App is running. Press Ctrl+C to exit.")

	// Block and wait for exit signal (with 5s timeout for cleanup)
	signailUtils.WaitExit(5 * time.Second)
}
```

### 3. Struct Utilities (`structUtils`)

Helpers for struct manipulation, deep empty checks, and map conversion.

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

	// Copy fields with same name and type
	structUtils.CopyIntersectionStruct(user, dto)
	fmt.Println(dto.Name) // Output: Alice

	// Check if struct is empty (deep check)
	emptyUser := &User{}
	fmt.Println(structUtils.IsStructEmpty(emptyUser)) // Output: true
}
```

### 4. Array Utilities (`arrayUtils`)

Common array/slice operations and a Set implementation.

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/arrayUtils"
)

func main() {
	// Check if item exists in slice
	nums := []int{1, 2, 3}
	if arrayUtils.InArray(2, nums) {
		fmt.Println("Found 2!")
	}

	// String Set
	set := arrayUtils.NewSet("a", "b")
	set.Add("c")
	if set.Has("a") {
		fmt.Println("Set has 'a'")
	}
}
```

### 5. Version Utilities (`versionUtils`)

Semantic versioning comparison helpers.

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/versionUtils"
)

func main() {
	v1 := "1.2.3"
	v2 := "1.3.0"

	if versionUtils.LessThan(v1, v2) {
		fmt.Printf("%s is older than %s\n", v1, v2)
	}
}
```

### 6. YAML Utilities (`yamlutil`)

Simple wrapper for reading and writing YAML files.

```go
package main

import "github.com/reggiepy/goutils/yamlutil"

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

### 7. Enum Utilities (`enumUtils`)

String-based enum validation.

```go
package main

import (
	"fmt"
	"github.com/reggiepy/goutils/enumUtils"
)

func main() {
	// Define allowed values
	status := enumUtils.NewEnum([]string{"pending", "active", "closed"}, "pending")

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
