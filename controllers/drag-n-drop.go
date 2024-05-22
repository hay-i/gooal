package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/chronologger/components"
	"github.com/labstack/echo/v4"
)

func Drag_n_Drop(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.Drag_n_Drop()

		return renderBase(c, component)
	}
}
