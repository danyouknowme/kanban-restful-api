package routes

import (
	"kanban/pkg/api"

	"github.com/gin-gonic/gin"
)

func TaskRoute(router *gin.Engine) {
	router.POST("/tasks", api.CreateTask())
	router.PUT("/tasks/list/same", api.EditTaskListSameColumn())
	router.PUT("/tasks/list/diff", api.EditTaskListDifferentColumn())
	router.PUT("/tasks/subtask", api.EditSubtaskStatus())
}
