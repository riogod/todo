package todo_list

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

func Setup(router *gin.RouterGroup, db *gorm.DB) {
	sevice := &Service{
		db,
	}

	todo := router.Group("/todo/list")
	{
		todo.GET(":id", get(sevice))
		todo.POST("", create(sevice))
		todo.PATCH(":id", func(ctx *gin.Context) {})
		todo.DELETE(":id", func(ctx *gin.Context) {})
	}
}

func get(s *Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		model, _ := s.Get(id)

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
