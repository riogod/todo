package apiv1

import (
	"fmt"
	"net/http"
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
		fmt.Println("Controller")
		responseMap := h.service.TodoList.GetAllRecords()
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body:    responseMap,
		})
	}
}

func (h *TodoListHandler) getById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
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
		id := ctx.Param("id")

		var requestParams RequestTodoListDTO
		err := ctx.BindJSON(&requestParams)
		if err != nil {
			fmt.Println("unable parse request json")
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
		id := ctx.Param("id")
		err := h.service.TodoList.Delete(id)
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
