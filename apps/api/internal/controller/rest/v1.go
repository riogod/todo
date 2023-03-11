package apiv1

import (
	"net/http"
	"todo_api/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRestHandlers(router *gin.Engine, service *service.Service) gin.RouterGroup {
	api := router.Group("/api")

	SetupTodoList(api, service)
	SetupTodoItem(api, service)

	return router.RouterGroup
}

func JSONError(ctx *gin.Context, code string, text string) {
	ctx.JSON(http.StatusOK, ResponseERROR_DTO{
		Success: false,
		Error: ResponseErrorDTO{
			ErrorCode:    code,
			ErrorMessage: text,
		},
	})
}
