package formparser

import (
	"net/url"

	"github.com/labstack/echo/v4"
)

func ParseForm(c echo.Context) (url.Values, error) {
	if err := c.Request().ParseForm(); err != nil {
		return nil, err
	}

	formValues, err := c.FormParams()
	if err != nil {
		return nil, err
	}

	return formValues, nil
}
