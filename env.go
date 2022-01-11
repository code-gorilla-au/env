package env

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	// FatalOnMissingEnv defines whether the library should fatal on encountering a missing environmental value
	FatalOnMissingEnv = false
)

// LoadEnvFile - loads .env file
func LoadEnvFile(path string) {
	if path == "" {
		log.Println("[env] No file provided, ")
		return
	}
	if err := godotenv.Overload(path); err != nil {
		log.Printf("[env] error: %s", err)
		return
	}
	log.Printf("[env] %s loaded", path)
}

// GetAsString reads an environment or returns a default value
func GetAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if FatalOnMissingEnv {
		log.Fatalf("[env] environment variable [%s] not set", key)
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
