package api

import (
	api_errors "hackthon/service/api/errors"
	"hackthon/service/models"
)

func (api *HackthonApiImpl) GetTodoTaskByIDAPI(taskId int64) (*models.TodoResponse, error) {
	todoTaskResp, err := api.db.GetTodoTaskByIDDB(taskId)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return todoTaskResp, nil
}
