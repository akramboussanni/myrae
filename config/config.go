package config

import (
	"encoding/base64"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecret []byte

func Init() {
	godotenv.Load()
	JwtSecret, err := base64.StdEncoding.DecodeString(os.Getenv("MYRAE_HMAC_SECRET"))
	if err != nil {
		panic(err)
	}
}
