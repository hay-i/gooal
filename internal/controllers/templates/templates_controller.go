package templates

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"
	"github.com/hay-i/gooal/internal/formparser"
	"github.com/hay-i/gooal/internal/models"

	"github.com/labstack/echo/v4"
)

func BuildGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		goal := c.QueryParam("goal")
		focus := c.QueryParam("focus")

		tokenString, err := auth.GetTokenFromCookie(c)
		if err != nil {
			return auth.HandleInvalidToken(c, "You must be logged in to access that page")
		}

		parsedToken, err := auth.ParseToken(tokenString)
		if err != nil {
			return auth.HandleInvalidToken(c, "Invalid or expired token")
		}

		username := auth.TokenToUsername(parsedToken)
		component := components.Build(goal, focus, username)

		return controllers.RenderNoBase(c, component)
	}
}

func InputGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Request().ParseForm(); err != nil {
			return err
		}

		var inputType models.QuestionType
		inputType = models.QuestionType(c.QueryParam("inputType"))
		order := c.QueryParam("order")

		objectId := primitive.NewObjectID()
		component := components.TemplateBuilderInput(inputType, objectId.Hex(), order)

		return controllers.RenderNoBase(c, component)
	}
}

func SavePOST(database *mongo.Database, client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		formValues, err := formparser.ValidateFormValues(c)
		if err != nil {
			return err
		}

		// TODO: Add validations for template builder
		models.Template{}.FromForm(formValues).Save(database)

		return controllers.RenderNoBase(c, components.Save("Template"))
	}
}

func InputDELETE() echo.HandlerFunc {
	return func(c echo.Context) error {
		return controllers.RenderNoBase(c, components.DeleteInput())
	}
}
