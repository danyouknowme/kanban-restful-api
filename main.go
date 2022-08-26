package main

import (
	"kanban/pkg/database"
	"kanban/pkg/routes"
	"kanban/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	config := utils.LoadConfig()

	database.ConnectDB()

	app.Use(cors.Default())

	routes.BoardRoute(app)
	routes.TaskRoute(app)

	app.Run(":" + config.Port)
}
