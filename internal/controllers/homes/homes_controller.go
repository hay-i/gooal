package homes

import (
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"
	"github.com/labstack/echo/v4"
)

func HomeGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.Home()

		return controllers.RenderBase(c, component)
	}
}
