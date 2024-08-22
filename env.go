package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	logPrefix = "[env]: "
)

var (
	// strictMode defines whether the library should fatal on encountering a missing environmental value
	strictMode = false
)

// // WithStrictMode defines whether the library should fatal on encountering a missing environmental value
func WithStrictMode() {
	strictMode = true
}

// LoadEnvFile - loads .env file
func LoadEnvFile(path string) bool {
	if path == "" {
		fmt.Printf("%s No path provided \n", logPrefix)
		return false
	}

	if err := godotenv.Overload(path); err != nil {
		fmt.Printf("%s error loading file [%s]: %s \n", logPrefix, path, err)
		return false
	}

	fmt.Printf("%s file [%s] loaded \n", logPrefix, path)
	return true
}

// GetAsString reads an environment or returns a default value
func GetAsString(key string) string {
	val, err := getString(key)
	if err != nil {
		panic(err)
	}

	return val
}

// GetAsStringWithDefault reads an environment variable or returns a default value.
// Ignore's strict mode.
func GetAsStringWithDefault(key string, defaultValue string) string {
	val, err := getString(key)
	if err != nil {
		return defaultValue
	}

	return val
}

// getString reads an environment variable, if strict mode is on, return an error.
func getString(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	if !strictMode {
		return "", nil
	}

	errResp := fmt.Errorf("%s environment variable [%s] not set", logPrefix, key)
	return "", errResp
}

// GetAsInt reads an environment variable into integer or returns a default value
func GetAsInt(name string) int {
	value, err := getInt(name)
	if err != nil {
		panic(err)
	}

	return value
}

// GetAsIntWithDefault reads an environment variable into an integer or returns a default value.
// Ignore's strict mode.
func GetAsIntWithDefault(name string, defaultValue int) int {
	value, err := getInt(name)
	if err != nil {
		return defaultValue
	}

	return value
}

// getInt reads an environment variable into an integer, if strict mode is on, return an error.
func getInt(name string) (int, error) {
	valueStr, err := getString(name)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil && strictMode {
		return 0, fmt.Errorf("%s environment variable [%s] not an int", logPrefix, name)
	}

	return value, nil
}

// GetAsBool reads an environment variable into a bool or return default value
func GetAsBool(name string) bool {
	val, err := getBool(name)
	if err != nil {
		panic(err)
	}

	return val
}

// GetAsBoolWithDefault reads an environment variable into a bool or returns a default value.
// Ignore's strict mode.
func GetAsBoolWithDefault(name string, defaultValue bool) bool {
	val, err := getBool(name)
	if err != nil {
		return defaultValue
	}

	return val
}

// getBool reads an environment variable into a bool, if strict mode is on, return an error.
func getBool(name string) (bool, error) {
	valStr, err := getString(name)
	if err != nil {
		return false, err
	}

	val, err := strconv.ParseBool(valStr)
	if err != nil && strictMode {
		return false, fmt.Errorf("%s environment variable [%s] not a boolean", logPrefix, name)
	}

	return val, nil
}

// GetAsSlice reads an environment variable into a string slice or returns the default value
func GetAsSlice(name string, sep string) []string {
	val, err := getSlice(name, sep)
	if err != nil {
		panic(err)
	}

	return val
}

// GetAsSliceWithDefault reads an environment variable into a string slice or returns a default value.
// Ignore's strict mode.
func GetAsSliceWithDefault(name string, sep string, defaultValue []string) []string {
	val, err := getSlice(name, sep)
	if err != nil {
		return defaultValue
	}

	return val
}

// getSlice reads an environment variable into a string slice, if strict mode is on, return an error.
func getSlice(name string, sep string) ([]string, error) {
	valStr, err := getString(name)
	if err != nil || len(valStr) == 0 {
		return []string{}, err
	}

	if sep == "" {
		sep = ","
	}

	return strings.Split(valStr, sep), nil
}
