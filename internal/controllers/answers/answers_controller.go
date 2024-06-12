package answers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"
	"github.com/hay-i/gooal/internal/form/parser"
	"github.com/hay-i/gooal/internal/form/validator"
	"github.com/hay-i/gooal/internal/models"
	"github.com/hay-i/gooal/internal/models/views"

	"github.com/labstack/echo/v4"
)

func AnswerTemplateGET(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		template := models.GetTemplate(database, id)

		questionViews := views.QuestionsToView(template.Questions)

		return controllers.RenderNoBase(c, components.Complete(template, questionViews))
	}
}

func AnswerTemplatePOST(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		template := models.GetTemplate(database, id)

		questionViews := views.QuestionsToView(template.Questions)

		formValues, err := parser.ParseForm(c)
		if err != nil {
			return err
		}

		questionViews = validator.ApplyAnsweringQuestionValidations(questionViews, formValues)

		if views.QuestionsHaveErrors(questionViews) {
			return controllers.RenderNoBase(c, components.Complete(template, questionViews))
		}

		tokenString, err := auth.GetTokenFromCookie(c)
		if err != nil {
			return auth.HandleInvalidToken(c, "You must be logged in to access that page")
		}

		parsedToken, err := auth.ParseToken(tokenString)
		if err != nil {
			return auth.HandleInvalidToken(c, "Invalid or expired token")
		}

		username := auth.TokenToUsername(parsedToken)
		questionAnswers := views.QuestionAnswersFromForm(questionViews)
		models.Answer{}.FromForm(template.ID, username, questionAnswers).Save(database)

		return controllers.RenderNoBase(c, components.Save("Response to template"))
	}
}
