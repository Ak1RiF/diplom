package helpers

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		//log
		return errors.New("No .env file found")
	}
	return nil
}

func GetByKey(key string) string {
	result, exists := os.LookupEnv(key)
	if exists != true {
		return "no value"
	}
	return result
}
