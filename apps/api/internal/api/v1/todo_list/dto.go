package todo_list

import "time"

type RequestTodoListDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ResponseTodoListItemDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

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
