package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"
	"github.com/labstack/echo/v4"
)

func Home(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.Home()

		return renderBase(c, component)
	}
}
