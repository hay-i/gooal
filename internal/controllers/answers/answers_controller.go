package answers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hay-i/gooal/internal/auth"
	"github.com/hay-i/gooal/internal/components"
	"github.com/hay-i/gooal/internal/controllers"
	"github.com/hay-i/gooal/internal/formparser"
	"github.com/hay-i/gooal/internal/models"

	"github.com/labstack/echo/v4"
)

func AnswerTemplateGET(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		template := models.GetTemplate(database, id)

		questionViews := models.QuestionsToView(template.Questions)

		return controllers.RenderNoBase(c, components.Complete(template, questionViews))
	}
}

func AnswerTemplatePOST(database *mongo.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		template := models.GetTemplate(database, id)

		questionViews := models.QuestionsToView(template.Questions)

		formValues, err := formparser.ParseForm(c)
		if err != nil {
			return err
		}

		questionViews = models.ApplyAnsweringQuestionValidations(questionViews, formValues)

		if models.QuestionsHaveErrors(questionViews) {
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
		models.Answer{}.FromForm(template.ID, username, questionViews).Save(database)

		return controllers.RenderNoBase(c, components.Save("Response to template"))
	}
}
