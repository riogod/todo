package apiv1

import (
	"todo_api/internal/api/v1/todo_list"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) gin.RouterGroup {
	api := router.Group("/api")

	todo_list.Setup(api)

	return router.RouterGroup
}
