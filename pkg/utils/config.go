package utils

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	MongoUri          string `mapstructure:"MONGO_URI"`
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func LoadConfig() {
	_ = godotenv.Load(".env")

	AppConfig.Port = os.Getenv("PORT")
	AppConfig.MongoUri = os.Getenv("MONGO_URI")

	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/google_callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
