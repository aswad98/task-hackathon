package db

import (
	db_errors "hackthon/service/api/errors"
	"hackthon/service/models"
)

func (db *HackthonDBImpl) GetAllTodoTaskDB() ([]*models.TodoResponse, error) {
	todoTaskResp := make([]*models.TodoResponse, 0)
	err := db.dbConn.Select(&todoTaskResp, `SELECT * FROM usertodotask;`)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return todoTaskResp, nil
}

func (db *HackthonDBImpl) GetTodoTaskByIDDB(taskId int64) (*models.TodoResponse, error) {
	todoTaskDB := models.NewTodoResponse()
	err := db.dbConn.Get(todoTaskDB, `SELECT * FROM usertodotask WHERE task_id=?;`, taskId)
	if err != nil {
		return nil, db_errors.NewError("no record found")
	}
	return todoTaskDB, nil
}
