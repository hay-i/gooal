package formparser

import (
	"net/url"
	"strings"
	"time"

	"github.com/hay-i/gooal/logger"
	"github.com/hay-i/gooal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TemplateFromForm(formValues url.Values) models.Template {
	template := models.Template{
		Title:       formValues.Get("title"),
		Description: formValues.Get("description"),
		CreatedAt:   time.Now(),
	}

	formValues.Del("title")
	formValues.Del("description")

	templatesQuestions := []models.Question{}

	for key, value := range formValues {
		inputLabel := value[0]
		splitKey := strings.Split(key, "-")
		inputType, id := splitKey[0], splitKey[1]

		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.LogError("Error:", err)
		}

		templatesQuestions = append(templatesQuestions, models.Question{
			ID:    objectID,
			Label: inputLabel,
			Type:  models.QuestionType(inputType),
		})
	}

	return template
}
