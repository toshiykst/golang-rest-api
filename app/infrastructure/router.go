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

	e.GET("/posts", func(c echo.Context) error { return postController.GetPosts(c) })
	e.GET("/posts/:id", func(c echo.Context) error { return postController.GetPost(c) })
	e.POST("/posts", func(c echo.Context) error { return postController.CreatePost(c) })
	e.PUT("/posts/:id", func(c echo.Context) error { return postController.UpdatePost(c) })
	e.DELETE("/posts/:id", func(c echo.Context) error { return postController.DeletePost(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
