package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	MongoUri string `mapstructure:"MONGO_URI"`
}

func LoadConfig() Config {
	_ = godotenv.Load(".env")
	config := Config{
		Port:     os.Getenv("PORT"),
		MongoUri: os.Getenv("MONGO_URI"),
	}
	return config
}
