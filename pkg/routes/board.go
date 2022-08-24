package routes

import (
	"kanban/pkg/api"

	"github.com/gin-gonic/gin"
)

func BoardRoute(router *gin.Engine) {
	router.GET("/boards/:userId", api.GetAllBoards())
}
