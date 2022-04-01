package api

import (
	"hackthon/service/db"
	"hackthon/service/models"
)

type HackthonApi interface {
	CreatetodotaskAPI(todoTaskData *models.Todo) (*models.TodoResponse, error)
	GetAllTodoTaskAPI() ([]*models.TodoResponse, error)
	GetTodoTaskByIDAPI(taskId int64) (*models.TodoResponse, error)
}
type HackthonApiImpl struct {
	db *db.HackthonDBImpl
}

func NewHackthonApiImpl() *HackthonApiImpl {
	dbImpl := db.NewHackthonImpl()
	return &HackthonApiImpl{
		db: dbImpl,
	}
}

var _ HackthonApi = &HackthonApiImpl{}
