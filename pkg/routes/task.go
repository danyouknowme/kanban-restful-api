package routes

import (
	"kanban/pkg/api"

	"github.com/gin-gonic/gin"
)

func TaskRoute(router *gin.Engine) {
	router.POST("/tasks", api.CreateTask())
}
