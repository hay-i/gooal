package templates

import (
	"fmt"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"
	"github.com/hay-i/gooal/internal/flash"
	"github.com/hay-i/gooal/internal/form/parser"
	"github.com/hay-i/gooal/internal/form/validator"
	"github.com/hay-i/gooal/internal/models"
	"github.com/hay-i/gooal/internal/models/views"
	"github.com/hay-i/gooal/pkg/logger"

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
		component := components.Build(goal, focus, username, views.TemplateView{})

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

		orderInt, err := strconv.Atoi(order)
		if err != nil {
			logger.LogError("Error:", err)
		}

		questionView := views.QuestionView{
			Question: models.Question{
				ID:    objectId,
				Type:  inputType,
				Order: orderInt,
			},
		}

		component := components.TemplateBuilderInput(questionView)

		return controllers.RenderNoBase(c, component)
	}
}

func SavePOST(database *mongo.Database, client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := auth.GetTokenFromCookie(c)
		if err != nil {
			return auth.HandleInvalidToken(c, "You must be logged in to access that page")
		}

		parsedToken, err := auth.ParseToken(tokenString)
		if err != nil {
			return auth.HandleInvalidToken(c, "Invalid or expired token")
		}

		username := auth.TokenToUsername(parsedToken)

		formValues, err := parser.ParseForm(c)
		if err != nil {
			return err
		}

		templateView := validator.ApplyTemplateBuilderValidations(formValues)

		if templateView.HasErrors() {
			component := components.BuildTemplateForm(username, templateView)
			return controllers.RenderNoBase(c, component)
		}

		id := models.Template{}.FromForm(formValues).Save(database)

		flash.Add(c, "You've successfully saved your template", flash.Success)

		c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/templates/%s/complete", id))
		return c.NoContent(http.StatusOK)
	}
}

func InputDELETE() echo.HandlerFunc {
	return func(c echo.Context) error {
		return controllers.RenderNoBase(c, components.DeleteInput())
	}
}
