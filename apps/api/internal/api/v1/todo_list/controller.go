package todo_list

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

type Success struct {
	success bool
}

func Setup(router *gin.RouterGroup, db *gorm.DB) {
	sevice := &Service{
		db,
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

func get(s *Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		model, _ := s.GetAll()

		ctx.JSON(http.StatusOK, model)
	}
}

func getById(s *Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		model, _ := s.GetByID(id)

		ctx.JSON(http.StatusOK, model)
	}
}

func create(s *Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var createParams model.ToDoItemList
		err := ctx.BindJSON(&createParams)
		if err != nil {
			log.Fatal(err)
		}
		s.Create(&createParams)
		ctx.JSON(http.StatusOK, createParams)
	}

}

func update(s *Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var updateParams model.ToDoItemList
		err := ctx.BindJSON(&updateParams)
		if err != nil {
			fmt.Println(err)
		}
		model, err := s.Update(id, &updateParams)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(http.StatusOK, model)
	}

}

func delete(s *Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := s.Delete(id)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(http.StatusOK, Success{
			success: true,
		})
	}

}
