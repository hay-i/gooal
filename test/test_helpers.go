package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func SetupEchoContext() echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}
