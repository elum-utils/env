package env

import (
    "os"
    "testing"
    "time"
)

// Test for retrieving a string environment variable
func TestGetEnvString(t *testing.T) {
    // Set a test environment variable
    os.Setenv("TEST_STRING", "test")
    defer os.Unsetenv("TEST_STRING")  // Ensure cleanup after the test

    // Retrieve the environment variable as a string
    got := GetEnvString("TEST_STRING", "default")
    want := "test"

    // Check if the retrieved value matches the expected value
    if got != want {
        t.Errorf("got %q; want %q", got, want)
    }
}

// Test for retrieving an array of strings from an environment variable
func TestGetEnvArrayString(t *testing.T) {
    os.Setenv("TEST_ARRAY_STRING", "test1,test2,test3")
    defer os.Unsetenv("TEST_ARRAY_STRING")  // Ensure cleanup after the test

    // Retrieve the environment variable as an array of strings
    got := GetEnvArrayString("TEST_ARRAY_STRING", ",", nil)
    want := []string{"test1", "test2", "test3"}

    // Check if the retrieved array matches the expected array
    for i, v := range got {
        if v != want[i] {
            t.Errorf("got %v; want %v", got, want)
            break
        }
    }
}

// Test for retrieving an integer environment variable
func TestGetEnvInt(t *testing.T) {
    os.Setenv("TEST_INT", "123")
    defer os.Unsetenv("TEST_INT")  // Ensure cleanup after the test

    // Retrieve the environment variable as an integer
    got := GetEnvInt("TEST_INT", 0)
    want := 123

    // Check if the retrieved value matches the expected value
    if got != want {
        t.Errorf("got %d; want %d", got, want)
    }
}

// Test for retrieving a duration environment variable
func TestGetEnvDuration(t *testing.T) {
    os.Setenv("TEST_DURATION", "1m")
    defer os.Unsetenv("TEST_DURATION")  // Ensure cleanup after the test

    // Retrieve the environment variable as a duration
    got := GetEnvDuration("TEST_DURATION", 0)
    want := time.Minute

    // Check if the retrieved duration matches the expected duration
    if got != want {
        t.Errorf("got %v; want %v", got, want)
    }
}

// Test for retrieving a boolean environment variable
func TestGetEnvBool(t *testing.T) {
    os.Setenv("TEST_BOOL", "true")
    defer os.Unsetenv("TEST_BOOL")  // Ensure cleanup after the test

    // Retrieve the environment variable as a boolean
    got := GetEnvBool("TEST_BOOL", false)
    want := true

    // Check if the retrieved value matches the expected value
    if got != want {
        t.Errorf("got %v; want %v", got, want)
    }
}

// Test for retrieving a float64 environment variable
func TestGetEnvFloat64(t *testing.T) {
    os.Setenv("TEST_FLOAT", "123.456")
    defer os.Unsetenv("TEST_FLOAT")  // Ensure cleanup after the test

    // Retrieve the environment variable as a float64
    got := GetEnvFloat64("TEST_FLOAT", 0.0)
    want := 123.456

    // Check if the retrieved value matches the expected value
    if got != want {
        t.Errorf("got %f; want %f", got, want)
    }
}

// Test for retrieving an array of integers from an environment variable
func TestGetEnvArrayInt(t *testing.T) {
    os.Setenv("TEST_ARRAY_INT", "1,2,3")
    defer os.Unsetenv("TEST_ARRAY_INT")  // Ensure cleanup after the test

    // Retrieve the environment variable as an array of integers
    got := GetEnvArrayInt("TEST_ARRAY_INT", ",", nil)
    want := []int{1, 2, 3}

    // Check if the retrieved array matches the expected array
    for i, v := range got {
        if v != want[i] {
            t.Errorf("got %v; want %v", got, want)
            break
        }
    }
}

// Test for retrieving an array of durations from an environment variable
func TestGetEnvArrayDuration(t *testing.T) {
    os.Setenv("TEST_ARRAY_DURATION", "1m,2m,3m")
    defer os.Unsetenv("TEST_ARRAY_DURATION")  // Ensure cleanup after the test

    // Retrieve the environment variable as an array of durations
    got := GetEnvArrayDuration("TEST_ARRAY_DURATION", ",", nil)
    want := []time.Duration{1 * time.Minute, 2 * time.Minute, 3 * time.Minute}

    // Check if the retrieved array matches the expected array
    for i, v := range got {
        if v != want[i] {
            t.Errorf("got %v; want %v", got, want)
            break
        }
    }
}

// Test for retrieving a map[string]string from an environment variable
func TestGetEnvMapStringString(t *testing.T) {
    os.Setenv("TEST_MAP", "key1:val1,key2:val2,key3:val3")
    defer os.Unsetenv("TEST_MAP")

    got := GetEnvMapStringString("TEST_MAP", ",", ":", nil)
    want := map[string]string{
        "key1": "val1",
        "key2": "val2",
        "key3": "val3",
    }

    for k, v := range want {
        if got[k] != v {
            t.Errorf("for key %q, got %q; want %q", k, got[k], v)
        }
    }

    // Test default return
    os.Unsetenv("TEST_MAP")
    def := map[string]string{"default": "value"}
    gotDef := GetEnvMapStringString("TEST_MAP", ",", ":", def)
    if gotDef["default"] != "value" {
        t.Errorf("expected default value to be returned, got %v", gotDef)
    }
}