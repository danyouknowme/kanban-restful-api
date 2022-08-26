package main

import (
	"kanban/pkg/database"
	"kanban/pkg/routes"
	"kanban/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := utils.LoadConfig()

	database.ConnectDB()

	routes.BoardRoute(router)
	routes.TaskRoute(router)

	router.Run(":" + config.Port)
}
