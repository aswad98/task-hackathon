package models

type Todo struct {
	Task_id     *int64  `json:"task_id" db:"task_id"`
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      *string `json:"status" db:"status"`
}

func NewTodo() *Todo {
	return &Todo{
		Task_id:     new(int64),
		Title:       new(string),
		Description: new(string),
		Status:      new(string),
	}
}
