package api

import (
	"todo_api/internal/api/v1/todo_item"
	"todo_api/internal/types"

	"github.com/gin-gonic/gin"
)

func InitRestHandlers(router *gin.Engine, service *types.Service) gin.RouterGroup {
	api := router.Group("/api")

	todo_item.Setup(api, service)

	return router.RouterGroup
}
