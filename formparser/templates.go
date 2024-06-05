package formparser

import (
	"net/url"
	"strconv"
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
		Username:    formValues.Get("username"),
		CreatedAt:   time.Now(),
	}

	formValues.Del("title")
	formValues.Del("description")
	formValues.Del("username")

	templatesQuestions := []models.Question{}

	for key, value := range formValues {
		inputLabel := value[0]
		splitKey := strings.Split(key, "-")
		inputType, id, order := splitKey[0], splitKey[1], splitKey[2]

		orderInt, err := strconv.Atoi(order)
		if err != nil {
			logger.LogError("Error:", err)
		}

		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.LogError("Error:", err)
		}

		templatesQuestions = append(templatesQuestions, models.Question{
			ID:    objectID,
			Label: inputLabel,
			Type:  models.QuestionType(inputType),
			Order: orderInt,
		})
	}

	template.Questions = templatesQuestions

	return template
}
