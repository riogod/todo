package apiv1

import (
	"fmt"
	"net/http"
	"strconv"
	"todo_api/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoItemHandler struct {
	service *service.Service
}

func SetupTodoItem(router *gin.RouterGroup, services *service.Service) {

	h := TodoItemHandler{
		service: services,
	}
	todo := router.Group("/todo")
	{
		todo.GET("/:id", h.getById())
		todo.POST("/:id", h.create())
		todo.PATCH(":id", h.update())
		todo.DELETE(":id", h.delete())
	}

}

func (h *TodoItemHandler) getById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "list id must be an number ")
			return
		}

		responseMap, model_err := h.service.TodoItem.GetByID(id)
		if model_err != nil {
			JSONError(ctx, "NOT_FOUND", fmt.Sprintf("not found item with id=%v", id))
			return
		}

		listResp := ResponseTodoListItemDTO{
			ID:          fmt.Sprintf("%d", responseMap.List.ID),
			Title:       responseMap.List.Title,
			Description: responseMap.List.Description,
			Status:      responseMap.List.Status,
			CreatedAt:   responseMap.List.CreatedAt,
			UpdatedAt:   responseMap.List.UpdatedAt,
		}

		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoItemDTO{
				ID:          fmt.Sprintf("%d", responseMap.ID),
				List:        listResp,
				Title:       responseMap.Title,
				Description: responseMap.Description,
				Status:      responseMap.Status,
				CreatedAt:   responseMap.CreatedAt,
				UpdatedAt:   responseMap.UpdatedAt,
			},
		})
	}
}

func (h *TodoItemHandler) create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idItem := ctx.Param("id")
		id, err := strconv.ParseUint(idItem, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "list id in path must be an number ")
			return
		}

		var requestParams RequestTodoItemDTO
		err_bind := ctx.BindJSON(&requestParams)

		if err_bind != nil {
			JSONError(ctx, "PARSE_JSON", "unable parse request json")
			return
		}

		resp, create_err := h.service.TodoItem.Create(id, requestParams.Title, requestParams.Description, requestParams.Status)
		if create_err != nil {
			JSONError(ctx, "UNABLE_TO_CREATE", fmt.Sprintf("%v", create_err))
			return
		}
		listResp := ResponseTodoListItemDTO{
			ID:          fmt.Sprintf("%d", resp.List.ID),
			Title:       resp.List.Title,
			Description: resp.List.Description,
			Status:      resp.List.Status,
			CreatedAt:   resp.List.CreatedAt,
			UpdatedAt:   resp.List.UpdatedAt,
		}

		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoItemDTO{
				ID:          fmt.Sprintf("%d", resp.ID),
				List:        listResp,
				Title:       resp.Title,
				Description: resp.Description,
				Status:      resp.Status,
				CreatedAt:   resp.CreatedAt,
				UpdatedAt:   resp.UpdatedAt,
			},
		})
	}

}

func (h *TodoItemHandler) update() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idItem := ctx.Param("id")
		id, err := strconv.ParseUint(idItem, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "id in path must be an number ")
			return
		}

		var requestParams RequestTodoItemDTO
		err_bind := ctx.BindJSON(&requestParams)

		if err_bind != nil {
			fmt.Println("unable parse request json")
		}

		resp, create_err := h.service.TodoItem.Update(id, requestParams.ListID, requestParams.Title, requestParams.Description, requestParams.Status)
		if create_err != nil {
			JSONError(ctx, "UNABLE_TO_UPDATE", fmt.Sprintf("%v", create_err))
			return
		}

		listResp := ResponseTodoListItemDTO{
			ID:          fmt.Sprintf("%d", resp.List.ID),
			Title:       resp.List.Title,
			Description: resp.List.Description,
			Status:      resp.List.Status,
			CreatedAt:   resp.List.CreatedAt,
			UpdatedAt:   resp.List.UpdatedAt,
		}

		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoItemDTO{
				ID:          fmt.Sprintf("%d", resp.ID),
				List:        listResp,
				Title:       resp.Title,
				Description: resp.Description,
				Status:      resp.Status,
				CreatedAt:   resp.CreatedAt,
				UpdatedAt:   resp.UpdatedAt,
			},
		})
	}

}

func (h *TodoItemHandler) delete() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idItem := ctx.Param("id")
		id, err := strconv.ParseUint(idItem, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "id in path must be an number ")
			return
		}
		err = h.service.TodoItem.Delete(id)
		if err != nil {
			JSONError(ctx, "CANNOT_DELETE", fmt.Sprintf("cannot delete item with id=%v", id))
			return
		}
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body:    nil,
		})
	}

}
