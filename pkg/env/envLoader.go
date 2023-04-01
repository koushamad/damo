package env

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load()
}

func Get(key string, defaultValue string) string {
	env := os.Getenv(key)

	if env == "" {
		return defaultValue
	}

	return env
}
