package main

import (
	"hackthon/service/server"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	server.NewServerImpl(e)
	e.Logger.Fatal(e.Start(":3030"))
}
