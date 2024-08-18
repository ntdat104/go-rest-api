package form

type TodoItemUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
}
