package todo_list

import (
	"fmt"
	"net/http"
	db_service "todo_api/internal/service/db"

	"github.com/gin-gonic/gin"
	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

func Setup(router *gin.RouterGroup, db *gorm.DB) {
	sevice := &db_service.Service{

		DB: db,
	}

	todo := router.Group("/todo/list")
	{
		todo.GET("", get(sevice))
		todo.GET(":id", getById(sevice))
		todo.POST("", create(sevice))
		todo.PATCH(":id", update(sevice))
		todo.DELETE(":id", delete(sevice))
	}
}

func get(s *db_service.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var resultMap []model.ToDoItemList
		var responseMap []ResponseTodoListItemDTO
		s.GetAll(&resultMap)

		for _, item := range resultMap {
			mappedItem := ResponseTodoListItemDTO{
				ID:          fmt.Sprintf("%d", item.ID),
				Title:       item.Title,
				Description: item.Description,
				Status:      item.Status,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
			}
			responseMap = append(responseMap, mappedItem)
		}

		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body:    responseMap,
		})
	}
}

func getById(s *db_service.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		model, err := s.GetByID(id)
		if err != nil {

			ctx.JSON(http.StatusOK, ResponseERROR_DTO{
				Success: false,
				Error: ResponseErrorDTO{
					ErrorCode:    "NOTHING_FOUND",
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

func create(s *db_service.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var createParams model.ToDoItemList
		var requestParams RequestTodoListDTO

		err := ctx.BindJSON(&requestParams)

		createParams.Title = requestParams.Title
		createParams.Description = requestParams.Description
		createParams.Status = requestParams.Status

		if err != nil {
			ctx.JSON(http.StatusOK, ResponseERROR_DTO{
				Success: false,
				Error: ResponseErrorDTO{
					ErrorCode:    "UNABLE_TO_CREATE",
					ErrorMessage: "Cannot insert new list item in db table",
				},
			})
			return
		}
		createParams.ID = 0
		s.Create(&createParams)
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoListItemDTO{
				ID:          fmt.Sprintf("%d", createParams.ID),
				Title:       createParams.Title,
				Description: createParams.Description,
				Status:      createParams.Status,
				CreatedAt:   createParams.CreatedAt,
				UpdatedAt:   createParams.UpdatedAt,
			},
		})
	}

}

func update(s *db_service.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var updateParams model.ToDoItemList
		var requestParams RequestTodoListDTO

		err := ctx.BindJSON(&requestParams)

		updateParams.Title = requestParams.Title
		updateParams.Description = requestParams.Description
		updateParams.Status = requestParams.Status

		if err != nil {
			fmt.Println(err)
		}
		err = s.Update(id, &updateParams)
		if err != nil {

			ctx.JSON(http.StatusOK, ResponseERROR_DTO{
				Success: false,
				Error: ResponseErrorDTO{
					ErrorCode:    "NOTHING_TO_UPDATE",
					ErrorMessage: fmt.Sprintf("not found item with id=%v for update", id),
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body: ResponseTodoListItemDTO{
				ID:          fmt.Sprintf("%d", updateParams.ID),
				Title:       updateParams.Title,
				Description: updateParams.Description,
				Status:      updateParams.Status,
				CreatedAt:   updateParams.CreatedAt,
				UpdatedAt:   updateParams.UpdatedAt,
			},
		})
	}

}

func delete(s *db_service.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := s.Delete(id)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(http.StatusOK, ResponseOK_DTO{
			Success: true,
			Body:    nil,
		})
	}

}
