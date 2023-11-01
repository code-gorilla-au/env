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
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if !strictMode {
		return ""
	}

	msg := fmt.Sprintf("%s environment variable [%s] not set \n", logPrefix, key)
	panic(msg)
}

// GetAsInt reads an environment variable into integer or returns a default value
func GetAsInt(name string) int {
	valueStr := GetAsString(name)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		msg := fmt.Sprintf("%s environment variable [%s] not an int \n", logPrefix, name)
		panic(msg)
	}
	return value
}

// GetAsBool reads an environment variable into a bool or return default value
func GetAsBool(name string) bool {
	valStr := GetAsString(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	if !strictMode {
		return false

	}

	msg := fmt.Sprintf("%s environment variable [%s] not a boolean \n", logPrefix, name)
	panic(msg)

}

// GetAsSlice reads an environment variable into a string slice or returns the default value
func GetAsSlice(name string, sep string) []string {
	valStr := GetAsString(name)

	if valStr == "" {
		return []string{}
	}

	if sep == "" {
		sep = ","
	}

	return strings.Split(valStr, sep)
}
