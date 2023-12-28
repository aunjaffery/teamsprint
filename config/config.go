package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type envVars struct {
	DNS       string
	JwtSecret string
}

var Envs envVars

func LoadConfig(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongouri := os.Getenv("MONGOURI")
	jwt_secret := os.Getenv("JWT_SECRET")
	Envs = envVars{
		DNS:       mongouri,
		JwtSecret: jwt_secret,
	}
}
