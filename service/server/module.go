package server

import (
	"hackthon/service/api"

	"github.com/labstack/echo"
)

type ServerImpl interface {
	MakeTodoTask(c echo.Context) error
	GetAllTodoTask(c echo.Context) error
	GetTodoTaskByID(c echo.Context) error
}
type Server struct {
	api *api.HackthonApiImpl
}

func NewServer() *Server {
	return &Server{
		api: api.NewHackthonApiImpl(),
	}
}
func NewServerImpl(e *echo.Echo) {
	server := NewServer()
	e.POST("/maketodotask", server.MakeTodoTask)
	e.GET("/getalltodotask", server.GetAllTodoTask)
	e.GET("/gettodotaskbyid", server.GetTodoTaskByID)
}

var _ ServerImpl = &Server{}
