package routes

import (
	"kanban/pkg/api"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.GET("/google/login", api.GoogleLogin())
	router.GET("/google/callback", api.GoogleCallback())
}
