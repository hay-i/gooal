package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/components"
	"github.com/hay-i/gooal/db"
	"github.com/hay-i/gooal/formparser"
	"github.com/hay-i/gooal/logger"
	"github.com/hay-i/gooal/models"

	"github.com/labstack/echo/v4"
)

func Build() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		goal := c.QueryParam("goal")
		focus := c.QueryParam("focus")

		component := components.Build(goal, focus)

		return renderNoBase(c, component)
	}
}

func Builder() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		var inputType models.QuestionType
		inputType = models.QuestionType(c.QueryParam("inputType"))

		objectId := primitive.NewObjectID()
		component := components.Builder(inputType, objectId.Hex())

		return renderNoBase(c, component)
	}
}

func Save(database *mongo.Database, client *mongo.Client, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		formValues, err := c.FormParams()
		if err != nil {
			return err
		}
		for key, values := range formValues {
			logger.LogInfo("Key: %s, Value: %s", key, values[0])
		}

		db.SaveTemplate(database, ctx, formparser.TemplateFromForm(formValues))

		return renderNoBase(c, components.Save())
	}
}

func Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.LogInfo("Delete")
		return renderNoBase(c, components.Delete())
	}
}
