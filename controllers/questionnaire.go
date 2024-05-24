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

		component := components.StepTwo(goal)
		return renderNoBase(c, component)
	}
}
