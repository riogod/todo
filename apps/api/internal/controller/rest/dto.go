package apiv1

import (
	"time"

	model "github.com/riogod/todo/libs/gomodels"
)

type ResponseErrorDTO struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type ResponseOK_DTO struct {
	Success bool        `json:"success"`
	Body    interface{} `json:"body,omitempty"`
}
type ResponseERROR_DTO struct {
	Success bool             `json:"success"`
	Error   ResponseErrorDTO `json:"error"`
}

type RequestTodoListDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type RequestTodoItemDTO struct {
	Title       string `json:"title"`
	ListID      string `json:"list_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ResponseTodoListItemDTO struct {
	ID          string            `json:"id"`
	Items       *[]model.ToDoItem `json:"items"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ResponseTodoItemDTO struct {
	ID          string                  `json:"id"`
	List        ResponseTodoListItemDTO `json:"list"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Status      string                  `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
