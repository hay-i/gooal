package models

import (
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hay-i/gooal/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionType string

const (
	TextQuestion     QuestionType = "text"
	NumberQuestion   QuestionType = "number"
	RangeQuestion    QuestionType = "range"
	SelectQuestion   QuestionType = "select"
	RadioQuestion    QuestionType = "radio"
	TextAreaQuestion QuestionType = "text_area"
	CheckboxQuestion QuestionType = "checkbox"
)

type Template struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	Username    string             `bson:"username"`
	Questions   []Question         `bson:"questions,omitempty"`
}

type Question struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Label   string             `bson:"label"`
	Type    QuestionType       `bson:"type"`
	Options []string           `bson:"options,omitempty"`
	Min     int                `bson:"min,omitempty"`
	Max     int                `bson:"max,omitempty"`
	Order   int                `bson:"order"`
}

type QuestionView struct {
	Question
	Error string `bson:"error,omitempty"`
	Value string `bson:"value,omitempty"`
}

func (t Template) FromForm(formValues url.Values) Template {
	t.Title = formValues.Get("title")
	t.Description = formValues.Get("description")
	t.Username = formValues.Get("username")
	t.CreatedAt = time.Now()

	formValues.Del("title")
	formValues.Del("description")
	formValues.Del("username")

	templatesQuestions := []Question{}

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

		templatesQuestions = append(templatesQuestions, Question{
			ID:    objectID,
			Label: inputLabel,
			Type:  QuestionType(inputType),
			Order: orderInt,
		})
	}

	t.Questions = templatesQuestions

	return t
}

func SortQuestionsByOrder(qs []QuestionView) []QuestionView {
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].Order < qs[j].Order
	})

	return qs
}
