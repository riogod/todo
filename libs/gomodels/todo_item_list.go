package model

import (
	"time"

	"gorm.io/gorm"
)

type TodoList struct {
	ID          uint64 `gorm:"primary_key;column:id;"`
	Title       string `gorm:"column:title;type:text"`
	Description string `gorm:"column:description;type:text"`
	Status      string `gorm:"column:status;type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t *TodoList) TableName() string {
	return "todo_item_list"
}
