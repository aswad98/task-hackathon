package api

import (
	api_errors "hackthon/service/api/errors"
	"hackthon/service/models"
)

func (api *HackthonApiImpl) GetAllTodoTaskAPI() ([]*models.TodoResponse, error) {

	complaintsResp, err := api.db.GetAllTodoTaskDB()
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return complaintsResp, nil
}
