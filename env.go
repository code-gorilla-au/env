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
func LoadEnvFile(path string) {
	if path == "" {
		fmt.Printf("%s No path provided", logPrefix)
		return
	}
	if err := godotenv.Overload(path); err != nil {
		fmt.Printf("%s error loading file [%s]: %s", logPrefix, path, err)
		return
	}
	fmt.Printf("%s file [%s] loaded", logPrefix, path)
}

// GetAsString reads an environment or returns a default value
func GetAsString(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if strictMode {
		msg := fmt.Sprintf("%s environment variable [%s] not set", logPrefix, key)
		panic(msg)
	}
	return ""
}

// GetAsInt reads an environment variable into integer or returns a default value
func GetAsInt(name string) int {
	valueStr := GetAsString(name)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		msg := fmt.Sprintf("%s environment variable [%s] not an int", logPrefix, name)
		panic(msg)
	}
	return value
}

// GetAsBool reads an environment variable into a bool or return default value
func GetAsBool(name string) bool {
	valStr := GetAsString(name)
	if val, err := strconv.ParseBool(valStr); err != nil {
		return val
	}
	if strictMode {
		msg := fmt.Sprintf("%s environment variable [%s] not a boolean", logPrefix, name)
		panic(msg)
	}
	return false
}

// GetAsSlice reads an environment variable into a string slice or returns the default value
func GetAsSlice(name string, sep string) []string {
	valStr := GetAsString(name)
	return strings.Split(valStr, sep)
}
