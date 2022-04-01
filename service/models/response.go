package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type TodoResponse struct {
	Task_id     *int64  `json:"task_id" db:"task_id"`
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      *string `json:"status" db:"status"`
}

func NewTodoResponse() *TodoResponse {
	return &TodoResponse{
		Task_id:     new(int64),
		Title:       new(string),
		Description: new(string),
		Status:      new(string),
	}
}
