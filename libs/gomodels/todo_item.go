package model

import (
	"time"

	"gorm.io/gorm"
)

type ToDoItem struct {
	ID          uint64 `gorm:"primary_key;column:id;"     json:"id"`
	Title       string `gorm:"column:title;type:text"                json:"title"`
	Description string `gorm:"column:description;type:text"          json:"description"`
	Status      string `gorm:"column:status;type:text"               json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t *ToDoItem) TableName() string {
	return "todo_item"
}
