package models

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hay-i/gooal/internal/db"
	"github.com/hay-i/gooal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Template struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	Username    string             `bson:"username"`
	Questions   []Question         `bson:"questions,omitempty"`
}

type TemplateView struct {
	Title            string
	TitleError       string
	Description      string
	DescriptionError string
	QuestionViews    []QuestionView
}

func (t Template) FromForm(formValues url.Values) Template {
	t.Title = formValues.Get("title")
	t.Description = formValues.Get("description")
	t.Username = formValues.Get("username")
	t.CreatedAt = time.Now()

	formValues.Del("title")
	formValues.Del("description")
	formValues.Del("username")

	t.Questions = QuestionsFromForm(formValues)

	return t
}

func (t Template) Save(database *mongo.Database) {
	db.Save(database, "templates", t)
}

func GetTemplate(database *mongo.Database, id string) Template {
	return db.Get(database, "templates", id).(Template)
}

func (t TemplateView) HasErrors() bool {
	if t.TitleError != "" || t.DescriptionError != "" {
		return true
	}

	if QuestionsHaveErrors(t.QuestionViews) {
		return true
	}

	return false
}

// Could live in a validators package
func ApplyTemplateBuilderValidations(formValues url.Values) TemplateView {
	template := TemplateView{}

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

	templateQuestions := make([]QuestionView, len(formValues))

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

		templateQuestions[i] = QuestionView{
			Question: Question{
				ID:    objectID,
				Type:  QuestionType(inputType),
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
