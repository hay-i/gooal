package questionnaires

import (
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"

	"github.com/labstack/echo/v4"
)

func StepOneGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.StepOne()
		return controllers.RenderBase(c, component)
	}
}

func StepTwoGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		goal := c.QueryParam("goal")

		var nextOptions []string

		switch goal {
		case "fitness":
			nextOptions = []string{"5k", "weightloss"}
		case "finance":
			nextOptions = []string{"savings", "investments"}
		case "career":
			nextOptions = []string{"promotion", "pay raise"}
		}

		component := components.StepTwo(goal, nextOptions)
		return controllers.RenderNoBase(c, component)
	}
}
