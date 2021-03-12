package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/toshiykst/golang-rest-api/app/handler"
	i "github.com/toshiykst/golang-rest-api/app/infrastructure"
)

func main() {
	e := echo.New()

	db := i.NewDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.HandlePostService(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
