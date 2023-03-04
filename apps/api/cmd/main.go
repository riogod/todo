package main

import (
	"fmt"

	model "github.com/riogod/todo/libs/gomodels"
)

func main() {
	a := model.ToDoItem{
		ID:          "a",
		Title:       "test",
		Description: "test desc",
		Status:      "no status",
	}
	fmt.Println("test", a)
}
