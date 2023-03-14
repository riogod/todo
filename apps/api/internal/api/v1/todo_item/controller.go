package todo_item

import (
	"net/http"
	"strconv"
	"todo_api/internal/middleware"
	service_error "todo_api/internal/service_errors"
	"todo_api/internal/types"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup, services *types.Service) {

	service := SetupService(services)

	todo := router.Group("/todo")
	{
		todo.GET("/:id", middleware.ValidateIdInQuery(), getById(service))
		todo.GET("/", search(service))
		todo.POST("", middleware.ValidatePOSTBody(RequestBodyCreateSchema), create(service))
		todo.PATCH(":id", middleware.ValidateIdInQuery(), middleware.ValidatePOSTBody(RequestBodyUpdateSchema), update(service))
		todo.DELETE(":id", middleware.ValidateIdInQuery(), delete(service))
	}

}

func getById(s *TodoItemService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id, _ := ctx.Get("paramID")

		response, okResponse := s.GetByID(id.(uint64))
		if okResponse != nil {
			ctx.Error(okResponse)
			return
		}

		ctx.JSON(http.StatusOK, ResponoseDTO{
			Success: true,
			Body:    *response,
		})
	}
}

func create(s *TodoItemService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		requestBody, _ := ctx.Get("requestBody")

		typedRequest := requestBody.(map[string]interface{})

		title := typedRequest["title"].(string)
		listId := typedRequest["list_id"].(string)
		description := typedRequest["description"].(string)
		status := typedRequest["status"].(string)

		list_id, err := strconv.ParseUint(listId, 10, 64)
		if err != nil {
			ctx.Error(service_error.ServiceError("MUST_BE_NUMBER", "list id must be a number"))
			return
		}

		response, okResponse := s.Create(
			list_id,
			title,
			description,
			status,
		)
		if okResponse != nil {
			ctx.Error(okResponse)
			return
		}

		ctx.JSON(http.StatusOK, ResponoseDTO{
			Success: true,
			Body:    response,
		})
	}
}

func update(s *TodoItemService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := ctx.Get("paramID")
		requestBody, _ := ctx.Get("requestBody")

		response, okResponse := s.Update(
			id.(uint64),
			requestBody.(map[string]interface{}),
		)
		if okResponse != nil {
			ctx.Error(okResponse)
			return
		}

		ctx.JSON(http.StatusOK, ResponoseDTO{
			Success: true,
			Body:    *response,
		})
	}
}

func delete(s *TodoItemService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := ctx.Get("paramID")
		response := s.Delete(id.(uint64))
		if response != nil {
			ctx.Error(response)
			return
		}

		ctx.JSON(http.StatusOK, ResponoseDTO{
			Success: true,
			Body:    nil,
		})

	}
}

func search(s *TodoItemService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		title := ctx.Query("title")
		response, okResponse := s.Search(title)
		if okResponse != nil {
			ctx.Error(okResponse)
			return
		}
		ctx.JSON(http.StatusOK, ResponoseDTO{
			Success: true,
			Body:    response,
		})
	}
}
