package model

import (
	"time"

	"gorm.io/gorm"
)

type TodoItem struct {
	ID          uint64 `gorm:"primary_key;column:id;"`
	ListID      uint64 `json:"list_id"`
	Title       string `gorm:"column:title;type:text"`
	Description string `gorm:"column:description;type:text"`
	Status      string `gorm:"column:status;type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t *TodoItem) TableName() string {
	return "todo_item"
}
