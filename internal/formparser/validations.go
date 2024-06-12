package formparser

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/hay-i/gooal/internal/models"
	"github.com/hay-i/gooal/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseForm(c echo.Context) (url.Values, error) {
	if err := c.Request().ParseForm(); err != nil {
		return nil, err
	}

	formValues, err := c.FormParams()
	if err != nil {
		return nil, err
	}

	return formValues, nil
}

func ValidateSubmission(formValues url.Values) models.TemplateView {
	template := models.TemplateView{}

	template.Title = formValues.Get("title")
	template.Description = formValues.Get("description")

	if formValues.Get("title") == "" {
		template.TitleError = "Title is required"
	}

	if formValues.Get("description") == "" {
		template.DescriptionError = "Description is required"
	}

	formValues.Del("title")
	formValues.Del("description")

	templateQuestions := make([]models.QuestionView, len(formValues))

	i := 0
	for key, value := range formValues {
		splitKey := strings.Split(key, "-")
		inputType, id, order := splitKey[0], splitKey[1], splitKey[2]
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.LogError("Error:", err)
		}

		orderInt, err := strconv.Atoi(order)
		if err != nil {
			logger.LogError("Error:", err)
		}

		templateQuestions[i] = models.QuestionView{
			Question: models.Question{
				ID:    objectID,
				Type:  models.QuestionType(inputType),
				Order: orderInt,
			},
			Value: value[0],
		}

		if value[0] == "" {
			templateQuestions[i].Error = "This field is required."
		}

		i++
	}

	template.QuestionViews = templateQuestions

	return template
}
