package validator

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/hay-i/gooal/internal/models"
	"github.com/hay-i/gooal/internal/models/views"
	"github.com/hay-i/gooal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ApplyTemplateBuilderValidations(formValues url.Values) views.TemplateView {
	template := views.TemplateView{}

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
	formValues.Del("username")

	templateQuestions := make([]views.QuestionView, len(formValues))

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

		templateQuestions[i] = views.QuestionView{
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

func ApplyAnsweringQuestionValidations(questionViews []views.QuestionView, formValues url.Values) []views.QuestionView {
	for i := range questionViews {
		val := formValues.Get(questionViews[i].ID.Hex())
		questionViews[i].Value = val

		if val == "" {
			questionViews[i].Error = "This field is required."
		}
	}

	return questionViews
}
