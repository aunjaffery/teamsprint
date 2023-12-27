package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVars struct {
	DNS string
}

var Envs EnvVars

func LoadConfig(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongouri := os.Getenv("MONGOURI")
	Envs = EnvVars{
		DNS: mongouri,
	}
}
