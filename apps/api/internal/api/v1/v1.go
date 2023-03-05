package apiv1

import (
	"todo_api/internal/api/v1/todo_list"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) gin.RouterGroup {
	api := router.Group("/api")

	todo_list.Setup(api, db)

	return router.RouterGroup
}
