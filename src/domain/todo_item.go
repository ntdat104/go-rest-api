package domain

import "time"

type TodoItem struct {
	Id          int        `json:"id" gorm:"colume:id"`
	Title       *string    `json:"title" gorm:"colume:title"`
	Description *string    `json:"description" gorm:"colume:description"`
	Status      *int       `json:"status" gorm:"colume:status"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"colume:created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"colume:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }
