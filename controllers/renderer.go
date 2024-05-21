package controllers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/hay-i/chronologger/auth"
	"github.com/hay-i/chronologger/components"
	"github.com/hay-i/chronologger/views"
	"github.com/labstack/echo/v4"
)

func renderNoBase(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func renderWithoutNav(c echo.Context, component templ.Component) error {
	base := components.BaseBody(views.GetFlashes(c), component)

	return base.Render(c.Request().Context(), c.Response().Writer)
}

func renderBase(c echo.Context, component templ.Component) error {
	flashes := views.GetFlashes(c)
	base := components.PageBase(flashes, auth.IsLoggedIn(c), component)

	return base.Render(c.Request().Context(), c.Response().Writer)
}

func redirect(c echo.Context, url string) error {
	// Old note: I wanted to send in http.StatusCreated, but it seems that the redirect doesn't work with that status code
	// See: https://github.com/labstack/echo/issues/229#issuecomment-1518502318
	return c.Redirect(http.StatusSeeOther, url)
}
