package controllers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func redirect(c echo.Context, url string) error {
	// Ideally I wanted to send in http.StatusCreated, but it seems that the redirect doesn't work with that status code
	// See: https://github.com/labstack/echo/issues/229#issuecomment-1518502318
	return c.Redirect(http.StatusFound, url)
}
