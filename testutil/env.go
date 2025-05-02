package testutil

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetRefreshToken() string {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")
	if refreshToken == "" {
		log.Println("PIXIV_REFRESH_TOKEN is not set")
	}
	return refreshToken
}
