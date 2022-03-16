package env

import (
	"log"
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

func WithStrictMode() {
	strictMode = true
}

// LoadEnvFile - loads .env file
func LoadEnvFile(path string) {
	if path == "" {
		log.Printf("%s No path provided", logPrefix)
		return
	}
	if err := godotenv.Overload(path); err != nil {
		log.Printf("%s error loading file [%s]: %s", logPrefix, path, err)
		return
	}
	log.Printf("%s file [%s] loaded", logPrefix, path)
}

// GetAsString reads an environment or returns a default value
func GetAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if strictMode {
		log.Fatalf("%s environment variable [%s] not set", logPrefix, key)
	}
	return defaultValue
}

// GetAsInt reads an environment variable into integer or returns a default value
func GetAsInt(name string, defaultValue int) int {
	valueStr := GetAsString(name, "")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetAsBool reads an environment variable into a bool or return default value
func GetAsBool(name string, defaultValue bool) bool {
	valStr := GetAsString(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultValue
}

// GetAsSlice reads an environment variable into a string slice or returns the default value
func GetAsSlice(name string, defaultValue []string, sep string) []string {
	valStr := GetAsString(name, "")

	if valStr == "" {
		return defaultValue
	}
	return strings.Split(valStr, sep)
}
