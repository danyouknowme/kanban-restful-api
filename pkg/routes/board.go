package routes

import (
	"kanban/pkg/api"

	"github.com/gin-gonic/gin"
)

func BoardRoute(router *gin.Engine) {
	router.GET("/boards", api.GetAllBoards())
	router.POST("/boards", api.CreateNewBoard())
	router.POST("/boards/task", api.CreateNewBoardTask())
}
