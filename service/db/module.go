package db

import (
	"fmt"
	"hackthon/service/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type HackthonDB interface {
	MakeAllTodoTaskDB(todoTaskData *models.Todo) (*models.TodoResponse, error)
	GetAllTodoTaskDB() ([]*models.TodoResponse, error)
	GetTodoTaskByIDDB(taskId int64) (*models.TodoResponse, error)
}
type HackthonDBImpl struct {
	dbConn *sqlx.DB
}

func NewHackthonImpl() *HackthonDBImpl {
	dbConn, err := sqlx.Connect("mysql", "database:Aswad_database@123@tcp(127.0.0.1:3306)/hackathonDB")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("db is connected")
	}

	return &HackthonDBImpl{
		dbConn: dbConn,
	}
}

var _ HackthonDB = &HackthonDBImpl{}
