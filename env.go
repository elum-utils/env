package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// envMap stores environment variables loaded from *.env files at runtime.
// Variables from the OS environment (os.Getenv) take precedence over these.
var envMap = make(map[string]string)

// init loads all environment variables from *.env files located in the same
// directory as the compiled binary. These variables are stored in memory
// (envMap) and are only used if the variable is not present in the system
// environment (os.Getenv). Variables are never written into the system
// environment to avoid exposure.
func init() {
	exePath, err := os.Executable()
	if err != nil {
		return
	}
	dir := filepath.Dir(exePath)

	// Discover all *.env files in the binary directory
	files, err := filepath.Glob(filepath.Join(dir, "*.env"))
	if err != nil {
		return
	}

	// Parse each .env file line-by-line
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			// Ignore empty lines and comments
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			// Parse key=value pairs
			kv := strings.SplitN(line, "=", 2)
			if len(kv) != 2 {
				continue
			}
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])

			// Only load the value if it's not already in the system environment
			if _, exists := os.LookupEnv(key); !exists {
				envMap[key] = val
			}
		}
	}
}

// GetEnvString retrieves an environment variable's value as a string.
// It first checks the OS environment, then loaded *.env files, and finally falls back to the default.
func GetEnvString(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	if val, ok := envMap[key]; ok {
		return val
	}
	return defaultValue
}

// GetEnvArrayString retrieves a string slice from a delimited environment variable or returns the default.
func GetEnvArrayString(key string, split string, defaultValue []string) []string {
	if val := GetEnvString(key, ""); val != "" {
		return strings.Split(val, split)
	}
	return defaultValue
}

// GetEnvInt retrieves an environment variable's value as an integer.
// Panics if the value exists but is not a valid integer.
func GetEnvInt(key string, defaultValue int) int {
	if val := GetEnvString(key, ""); val != "" {
		intValue, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid integer: %v", key, err))
		}
		return intValue
	}
	return defaultValue
}

// GetEnvDuration retrieves an environment variable's value as a time.Duration.
// Panics if the value exists but is not a valid duration.
func GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if val := GetEnvString(key, ""); val != "" {
		durationValue, err := time.ParseDuration(val)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid duration: %v", key, err))
		}
		return durationValue
	}
	return defaultValue
}

// GetEnvBool retrieves an environment variable's value as a boolean.
// Panics if the value exists but is not a valid boolean.
func GetEnvBool(key string, defaultValue bool) bool {
	if val := GetEnvString(key, ""); val != "" {
		boolValue, err := strconv.ParseBool(val)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid boolean: %v", key, err))
		}
		return boolValue
	}
	return defaultValue
}

// GetEnvFloat64 retrieves an environment variable's value as a float64.
// Panics if the value exists but is not a valid float64.
func GetEnvFloat64(key string, defaultValue float64) float64 {
	if val := GetEnvString(key, ""); val != "" {
		floatValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid float64: %v", key, err))
		}
		return floatValue
	}
	return defaultValue
}

// GetEnvArrayInt retrieves an environment variable's value as a slice of integers.
// Panics if any value in the slice is not a valid integer.
func GetEnvArrayInt(key string, split string, defaultValue []int) []int {
	if val := GetEnvString(key, ""); val != "" {
		stringValues := strings.Split(val, split)
		intValues := make([]int, 0, len(stringValues))
		for _, str := range stringValues {
			intValue, err := strconv.Atoi(str)
			if err != nil {
				panic(fmt.Sprintf("Environment variable %s array contains an invalid integer: %s", key, str))
			}
			intValues = append(intValues, intValue)
		}
		return intValues
	}
	return defaultValue
}

// GetEnvArrayDuration retrieves an environment variable's value as a slice of time.Duration values.
// Panics if any value in the slice is not a valid duration.
func GetEnvArrayDuration(key string, split string, defaultValue []time.Duration) []time.Duration {
	if val := GetEnvString(key, ""); val != "" {
		stringValues := strings.Split(val, split)
		durationValues := make([]time.Duration, 0, len(stringValues))
		for _, str := range stringValues {
			durationValue, err := time.ParseDuration(str)
			if err != nil {
				panic(fmt.Sprintf("Environment variable %s array contains an invalid duration: %s", key, str))
			}
			durationValues = append(durationValues, durationValue)
		}
		return durationValues
	}
	return defaultValue
}

// GetEnvMapStringString retrieves an environment variable as a map[string]string.
// The variable should contain key-value pairs delimited by entryDelimiter and kvDelimiter.
// Panics if any entry doesn't contain exactly one key-value delimiter.
func GetEnvMapStringString(key string, entryDelimiter string, kvDelimiter string, defaultValue map[string]string) map[string]string {
	if val := GetEnvString(key, ""); val != "" {
		result := make(map[string]string)
		entries := strings.Split(val, entryDelimiter)
		for _, entry := range entries {
			kv := strings.SplitN(entry, kvDelimiter, 2)
			if len(kv) != 2 {
				panic(fmt.Sprintf("Environment variable %s contains invalid map entry: %s", key, entry))
			}
			result[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
		return result
	}
	return defaultValue
}
