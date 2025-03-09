# env Package

The `env` package provides utility functions to simplify retrieving environment variables in various data types within Go applications. It ensures a default value is returned if the environment variable is not set or invalid, with optional panic mechanisms for certain conversions.

## Installation

To use the `env` package, you need to include it in your project. Assuming it is part of your Go application's module, you can import it directly:

```go
import "github.com/elum-utils/env"
```

## Functions

### GetEnvString

```go
func GetEnvString(key, defaultValue string) string
```

Retrieves an environment variable's value as a string. Returns the `defaultValue` if the variable is not set.

### GetEnvArrayString

```go
func GetEnvArrayString(key string, split string, defaultValue []string) []string
```

Retrieves an environment variable's value as a slice of strings by splitting it using the provided delimiter. Returns the `defaultValue` if the variable is not set.

### GetEnvInt

```go
func GetEnvInt(key string, defaultValue int) int
```

Retrieves an environment variable's value as an integer. Returns the `defaultValue` if the variable is not set. Panics if the value exists and cannot be converted to an integer.

### GetEnvDuration

```go
func GetEnvDuration(key string, defaultValue time.Duration) time.Duration
```

Retrieves an environment variable's value as a `time.Duration`. Returns the `defaultValue` if the variable is not set. Panics if the value exists and cannot be converted to a duration.

### GetEnvBool

```go
func GetEnvBool(key string, defaultValue bool) bool
```

Retrieves an environment variable's value as a boolean. Returns the `defaultValue` if the variable is not set. Panics if the value exists and cannot be converted to a boolean.

### GetEnvFloat64

```go
func GetEnvFloat64(key string, defaultValue float64) float64
```

Retrieves an environment variable's value as a `float64`. Returns the `defaultValue` if the variable is not set. Panics if the value exists and cannot be converted to a float64.

### GetEnvArrayInt

```go
func GetEnvArrayInt(key string, split string, defaultValue []int) []int
```

Retrieves an environment variable's value as a slice of integers by splitting it using the provided delimiter. Returns the `defaultValue` if the variable is not set. Panics if any value in the slice cannot be converted to an integer.

### GetEnvArrayDuration

```go
func GetEnvArrayDuration(key string, split string, defaultValue []time.Duration) []time.Duration
```

Retrieves an environment variable's value as a slice of `time.Duration` by splitting it using the provided delimiter. Returns the `defaultValue` if the variable is not set. Panics if any value in the slice cannot be converted to a duration.

## Example Usage

Here is how you might use the `env` package in a Go application:

```go
package main

import (
    "fmt"
    "your_module_name/env"
    "time"
)

var (
    myString = env.GetEnvString("MY_STRING", "defaultString")
    myInt = env.GetEnvInt("MY_INT", 42)
    myDuration = env.GetEnvDuration("MY_DURATION", 30*time.Second)
    myBool = env.GetEnvBool("MY_BOOL", false)
    myFloat = env.GetEnvFloat64("MY_FLOAT", 3.14)
)

func main() {


    fmt.Println("String:", myString)
    fmt.Println("Int:", myInt)
    fmt.Println("Duration:", myDuration)
    fmt.Println("Bool:", myBool)
    fmt.Println("Float:", myFloat)
}
```

## License

This package is open source and available under the MIT License. Feel free to use and modify it as needed.