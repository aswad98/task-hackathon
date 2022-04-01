package server

import (
	"hackthon/service/models"
	"net/http"

	server_errors "hackthon/service/api/errors"

	"github.com/labstack/echo"
)

func (s *Server) MakeTodoTask(c echo.Context) error {
	todoTaskData := models.NewTodo()
	if err := c.Bind(&todoTaskData); err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	todoTaskResp, err := s.api.CreatetodotaskAPI(todoTaskData)
	if err != nil {
		return server_errors.NewUnauthorizedError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Your task has been added",
		Data:    todoTaskResp,
	}
	return c.JSON(http.StatusOK, result)
}
