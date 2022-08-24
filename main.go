package main

import (
	"os"

	"kanban/pkg/database"
	"kanban/pkg/routes"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	_ = godotenv.Load(".env")

	mongoUri := os.Getenv("MONGO_URI")
	database.ConnectDB(mongoUri)

	routes.BoardRoute(router)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
