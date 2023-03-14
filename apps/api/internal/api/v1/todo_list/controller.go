package todo_list

import (
	"net/http"
	"todo_api/internal/middleware"
	"todo_api/internal/types"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup, services *types.Service) {

	service := SetupService(services)

	todo := router.Group("/todo/list")
	{
		todo.GET("/:id", middleware.ValidateIdInQuery(), getById(service))
		todo.GET("/all", getAllWithItems(service))
		todo.GET("", getAll(service))
		todo.POST("", middleware.ValidatePOSTBody(RequestBodyCreateSchema), create(service))
		todo.PATCH(":id", middleware.ValidateIdInQuery(), middleware.ValidatePOSTBody(RequestBodyUpdateSchema), update(service))
		todo.DELETE(":id", middleware.ValidateIdInQuery(), delete(service))
	}

}

func getById(s *TodoListService) func(ctx *gin.Context) {
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

func getAll(s *TodoListService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		response, okResponse := s.GetAll()
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

func getAllWithItems(s *TodoListService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		response, okResponse := s.GetAllWithItems()
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

func create(s *TodoListService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		requestBody, _ := ctx.Get("requestBody")

		typedRequest := requestBody.(map[string]interface{})

		title := typedRequest["title"].(string)
		description := typedRequest["description"].(string)
		status := typedRequest["status"].(string)

		response, okResponse := s.Create(
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

func update(s *TodoListService) func(ctx *gin.Context) {
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

func delete(s *TodoListService) func(ctx *gin.Context) {
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
