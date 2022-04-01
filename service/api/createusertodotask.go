package api

import (
	api_errors "hackthon/service/api/errors"
	"hackthon/service/models"
)

func (api *HackthonApiImpl) CreatetodotaskAPI(todoTaskData *models.Todo) (*models.TodoResponse, error) {

	todoTaskResp, err := api.db.MakeAllTodoTaskDB(todoTaskData)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return todoTaskResp, nil
}
