package controllers

import (
	"github.com/a-h/templ"
	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/flash"
	"github.com/labstack/echo/v4"
)

func RenderNoBase(c echo.Context, component templ.Component) error {
	if c.Request().Header.Get("HX-Request") != "" {
		return component.Render(c.Request().Context(), c.Response().Writer)
	}
	return RenderBase(c, component)
}

func RenderWithoutNav(c echo.Context, component templ.Component) error {
	base := components.BaseBody(flash.Get(c), component)

	return base.Render(c.Request().Context(), c.Response().Writer)
}

func RenderBase(c echo.Context, component templ.Component) error {
	flashes := flash.Get(c)
	base := components.PageBase(flashes, auth.IsLoggedIn(c), component)

	return base.Render(c.Request().Context(), c.Response().Writer)
}
