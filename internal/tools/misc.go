package tools

import (
	"os"

	"github.com/joho/godotenv"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	PanicIfErr(err)

	return os.Getenv(key)
}
