package db

import (
	db_errors "hackthon/service/api/errors"
	"hackthon/service/models"
)

func (db *HackthonDBImpl) MakeAllTodoTaskDB(todoTaskData *models.Todo) (*models.TodoResponse, error) {
	todoTaskResp := models.NewTodoResponse()
	tx := db.dbConn.MustBegin()
	_, err := tx.NamedQuery(`INSERT INTO usertodotask(task_id,title,description,status)
		 VALUES(:task_id,:title,:description,:status);`, todoTaskData)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	err = db.dbConn.Get(todoTaskResp, `SELECT * FROM usertodotask WHERE task_id=?;`, *todoTaskData.Task_id)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return todoTaskResp, nil
}
