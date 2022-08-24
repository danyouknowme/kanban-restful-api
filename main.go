package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	_ = godotenv.Load(".env")

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
