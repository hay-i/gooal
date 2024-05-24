package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"

	"github.com/labstack/echo/v4"
)

func StepOne(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.StepOne()
		return renderBase(c, component)
	}
}

func StepTwo(database *mongo.Database) echo.HandlerFunc {
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
		return renderNoBase(c, component)
	}
}
