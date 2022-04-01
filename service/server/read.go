package server

import (
	server_errors "hackthon/service/api/errors"
	"hackthon/service/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (s *Server) GetAllTodoTask(c echo.Context) error {
	getTodoTaskResp, err := s.api.GetAllTodoTaskAPI()
	if err != nil {
		return server_errors.NewUnauthorizedError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Successfully fetched all the todo tasks",
		Data:    getTodoTaskResp,
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) GetTodoTaskByID(c echo.Context) error {
	taskIdStr := c.Param("task_id")
	taskId, err := strconv.ParseInt(taskIdStr, 0, 64)
	if err != nil {
		server_errors.NewInternalServerError(err.Error())
	}
	todoTaskResp, err := s.api.GetTodoTaskByIDAPI(taskId)
	if err != nil {
		return server_errors.NewUnauthorizedError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Successfully fetched the task",
		Data:    todoTaskResp,
	}
	return c.JSON(http.StatusOK, result)
}
