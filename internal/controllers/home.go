package controllers

import (
	"github.com/hay-i/gooal/internal/components"
	"github.com/labstack/echo/v4"
)

func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.Home()

		return renderBase(c, component)
	}
}
