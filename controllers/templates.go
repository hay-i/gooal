package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/auth"
	"github.com/hay-i/gooal/components"
	"github.com/hay-i/gooal/db"
	"github.com/hay-i/gooal/formparser"
	"github.com/hay-i/gooal/models"
	"github.com/hay-i/gooal/views"

	"github.com/labstack/echo/v4"
)

func Build() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		goal := c.QueryParam("goal")
		focus := c.QueryParam("focus")

		// TODO: Extract all this
		cookie, err := c.Cookie("token")
		if err != nil {
			views.AddFlash(c, "You must be logged in to access that page", views.FlashError)

			return redirect(c, "/login")
		}

		tokenString := cookie.Value

		parsedToken, err := auth.ParseToken(tokenString)

		if err != nil {
			views.AddFlash(c, "Invalid or expired token", views.FlashError)

			return redirect(c, "/login")
		}

		component := components.Build(goal, focus, parsedToken["sub"].(string))

		return renderNoBase(c, component)
	}
}

func Input() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		var inputType models.QuestionType
		inputType = models.QuestionType(c.QueryParam("inputType"))
		order := c.QueryParam("order")

		objectId := primitive.NewObjectID()
		component := components.TemplateBuilderInput(inputType, objectId.Hex(), order)

		return renderNoBase(c, component)
	}
}

func Save(database *mongo.Database, client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		formValues, err := formparser.ValidateFormValues(c)
		if err != nil {
			return err
		}

		// TODO: Add validations for template builder
		db.SaveTemplate(database, models.Template{}.FromForm(formValues))

		return renderNoBase(c, components.Save("Template"))
	}
}

func DeleteInput() echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderNoBase(c, components.DeleteInput())
	}
}

func CompleteTemplate(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		template := db.GetTemplate(database, id)

		questionViews := formparser.QuestionsToView(template.Questions)

		return renderNoBase(c, components.Complete(template, questionViews))
	}
}

func Complete(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		template := db.GetTemplate(database, id)

		questionViews := formparser.QuestionsToView(template.Questions)

		formValues, err := formparser.ValidateFormValues(c)
		if err != nil {
			return err
		}

		questionViews = formparser.ApplyValidations(questionViews, formValues)

		if formparser.HasErrors(questionViews) {
			return renderNoBase(c, components.Complete(template, questionViews))
		}

		// TODO: Save
		return renderNoBase(c, components.Save("Response to template"))
	}
}
