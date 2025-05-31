package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(envVar string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(envVar)
}
