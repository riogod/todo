package apiv1

import (
	"fmt"
	"net/http"
	"strconv"
	"todo_api/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoListHandler struct {
	service *service.Service
}

func SetupTodoList(router *gin.RouterGroup, services *service.Service) {

	h := TodoListHandler{
		service: services,
	}
	todo := router.Group("/todo/list")
	{
		todo.GET("", h.get())
		todo.GET(":id", h.getById())
		todo.POST("", h.create())
		todo.PATCH(":id", h.update())
		todo.DELETE(":id", h.delete())
	}

}

func (h *TodoListHandler) get() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		responseMap := h.service.TodoList.GetAllRecords()
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body:    responseMap,
		})
	}
}

func (h *TodoListHandler) getById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idItem := ctx.Param("id")
		id, err := strconv.ParseUint(idItem, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "list id in path must be an number ")
			return
		}

		model, err := h.service.TodoList.GetByID(id)
		if err != nil {

			ctx.JSON(http.StatusOK, ResponseERROR_DTO{
				Success: false,
				Error: ResponseErrorDTO{
					ErrorCode:    "NOT_FOUND",
					ErrorMessage: fmt.Sprintf("not found item with id=%v", id),
				},
			})
			return
		}

		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoListItemDTO{
				ID:          fmt.Sprintf("%d", model.ID),
				Title:       model.Title,
				Items:       model.Items,
				Description: model.Description,
				Status:      model.Status,
				CreatedAt:   model.CreatedAt,
				UpdatedAt:   model.UpdatedAt,
			},
		})
	}
}

func (h *TodoListHandler) create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestParams RequestTodoListDTO
		err := ctx.BindJSON(&requestParams)

		if err != nil {
			fmt.Println("unable parse request json")
		}
		resp, create_err := h.service.TodoList.Create(requestParams.Title, requestParams.Description, requestParams.Status)
		if create_err != nil {
			JSONError(ctx, "UNABLE_TO_CREATE", "Cannot insert new list item in db table")
			return
		}

		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoListItemDTO{
				ID:          fmt.Sprintf("%d", resp.ID),
				Title:       resp.Title,
				Description: resp.Description,
				Status:      resp.Status,
				CreatedAt:   resp.CreatedAt,
				UpdatedAt:   resp.UpdatedAt,
			},
		})
	}

}

func (h *TodoListHandler) update() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idItem := ctx.Param("id")
		id, err := strconv.ParseUint(idItem, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "list id in path must be an number ")
			return
		}

		var requestParams RequestTodoListDTO
		err = ctx.BindJSON(&requestParams)
		if err != nil {
			JSONError(ctx, "PARSE_JSON", "unable parse request json")
			return
		}

		resp, err := h.service.TodoList.Update(id, requestParams.Title, requestParams.Description, requestParams.Status)
		if err != nil {
			JSONError(ctx, "NOTHING_TO_UPDATE", fmt.Sprintf("not found item with id=%v for update", id))

			return
		}
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoListItemDTO{
				ID:          fmt.Sprintf("%d", resp.ID),
				Title:       resp.Title,
				Description: resp.Description,
				Status:      resp.Status,
				CreatedAt:   resp.CreatedAt,
				UpdatedAt:   resp.UpdatedAt,
			},
		})
	}

}

func (h *TodoListHandler) delete() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idItem := ctx.Param("id")
		id, err := strconv.ParseUint(idItem, 10, 64)
		if err != nil {
			JSONError(ctx, "MUST_BE_INT", "list id in path must be an number ")
			return
		}
		err = h.service.TodoList.Delete(id)
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
