package router

import (
	"clean-architecture/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//NewRouter creates a new router instance
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	return e
}
