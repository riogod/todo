package todo_list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/riogod/todo/libs/gomodels"
)

func Setup(router *gin.RouterGroup) {

	todo := router.Group("/todo")
	{
		todo.GET("", func(ctx *gin.Context) {
			ret := model.ToDoItemList{
				ID:          "asdas",
				Title:       "asdazxzxc",
				Description: "zzzz",
				Status:      "OK",
			}
			ctx.JSON(http.StatusOK, &ret)
		})
		todo.PUT("", func(ctx *gin.Context) {})

		todo.DELETE("", func(ctx *gin.Context) {})
	}
}

// func Get() func(ctx *gin.Context) {

// 	ret := model.ToDoItemList{
// 		ID:          "asdas",
// 		Title:       "asdazxzxc",
// 		Description: "zzzz",
// 		Status:      "OK",
// 	}
// 	return func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, &ret)
// 	}

// }
