package form

type TodoItemCreate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
}
