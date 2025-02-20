package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func RunChecks() {
	godotenv.Load(".env")
	checkEnv("APP_URL")
}

func checkEnv(key string) {
	if os.Getenv(key) == "" {
		panic(fmt.Sprintf("Environment variable %s is not set", key))
	}
}
