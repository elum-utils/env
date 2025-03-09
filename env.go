package env

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

// GetEnvString retrieves an environment variable's value as a string or returns the default value if not set.
func GetEnvString(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

// GetEnvArrayString retrieves an environment variable's value as a slice of strings or returns the default value if not set.
func GetEnvArrayString(key string, split string, defaultValue []string) []string {
    if value, exists := os.LookupEnv(key); exists {
        return strings.Split(value, split)
    }
    return defaultValue
}

// GetEnvInt retrieves an environment variable's value as an integer or returns the default value if not set.
// Panics if the value exists but cannot be converted to an integer.
func GetEnvInt(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists {
        intValue, err := strconv.Atoi(value)
        if err != nil {
            panic(fmt.Sprintf("Environment variable %s is not a valid integer: %v", key, err))
        }
        return intValue
    }
    return defaultValue
}

// GetEnvDuration retrieves an environment variable's value as a duration or returns the default value if not set.
// Panics if the value exists but cannot be converted to a duration.
func GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
    if value, exists := os.LookupEnv(key); exists {
        durationValue, err := time.ParseDuration(value)
        if err != nil {
            panic(fmt.Sprintf("Environment variable %s is not a valid duration: %v", key, err))
        }
        return durationValue
    }
    return defaultValue
}

// GetEnvBool retrieves an environment variable's value as a boolean or returns the default value if not set.
// Panics if the value exists but cannot be converted to a boolean.
func GetEnvBool(key string, defaultValue bool) bool {
    if value, exists := os.LookupEnv(key); exists {
        boolValue, err := strconv.ParseBool(value)
        if err != nil {
            panic(fmt.Sprintf("Environment variable %s is not a valid boolean: %v", key, err))
        }
        return boolValue
    }
    return defaultValue
}

// GetEnvFloat64 retrieves an environment variable's value as a float64 or returns the default value if not set.
// Panics if the value exists but cannot be converted to a float64.
func GetEnvFloat64(key string, defaultValue float64) float64 {
    if value, exists := os.LookupEnv(key); exists {
        floatValue, err := strconv.ParseFloat(value, 64)
        if err != nil {
            panic(fmt.Sprintf("Environment variable %s is not a valid float64: %v", key, err))
        }
        return floatValue
    }
    return defaultValue
}

// GetEnvArrayInt retrieves an environment variable's value as a slice of integers or returns the default value if not set.
// Panics if any value in the slice cannot be converted to an integer.
func GetEnvArrayInt(key string, split string, defaultValue []int) []int {
    if value, exists := os.LookupEnv(key); exists {
        stringValues := strings.Split(value, split)
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

// GetEnvArrayDuration retrieves an environment variable's value as a slice of durations or returns the default value if not set.
// Panics if any value in the slice cannot be converted to a duration.
func GetEnvArrayDuration(key string, split string, defaultValue []time.Duration) []time.Duration {
    if value, exists := os.LookupEnv(key); exists {
        stringValues := strings.Split(value, split)
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