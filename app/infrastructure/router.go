package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/toshiykst/golang-rest-api/app/interfaces/controllers"
)

// RunRouter runs router.
func RunRouter() {
	e := echo.New()
	db := NewDB()

	postController := controllers.NewPostController(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/posts", func(c echo.Context) error { return postController.CreatePost(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
