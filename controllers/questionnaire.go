package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"

	"github.com/labstack/echo/v4"
)

func GetStarted(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.GetStarted()
		return renderBase(c, component)
	}
}

func StepOne(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		selectedOptions := c.Request().Form["options"]

		component := components.StepTwo(selectedOptions)
		return renderNoBase(c, component)
	}
}
