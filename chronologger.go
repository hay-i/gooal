package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
    component := hello("World")
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return component.Render(c.Request().Context(), c.Response().Writer)
    })
	e.Logger.Fatal(e.Start(":1323"))
}
