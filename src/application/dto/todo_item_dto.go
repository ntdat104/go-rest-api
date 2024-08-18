package dto

type TodoItemDto struct {
	Id          int     `json:"id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *int    `json:"status,omitempty"`
}
