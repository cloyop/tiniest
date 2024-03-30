package config

import (
	"log"
	"os"

	"github.com/cloyop/tiniest/internal/utils"
	"github.com/joho/godotenv"
)

func LoadEnvConfig() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file", err)
		}

	}
}
func LoadHelpers() {
	utils.DeletePublicExceed()
	utils.UpdateMaxLinks()
	utils.RunIpQuequeChannel()
	utils.CleanEmailsInterval()
}
