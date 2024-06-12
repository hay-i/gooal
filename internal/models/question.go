package models

import (
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/hay-i/gooal/pkg/logger"
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
	Error string
	Value string
}

func (q Question) OrderToString() string {
	return strconv.Itoa(q.Order)
}

func SortQuestionsByOrder(qs []QuestionView) []QuestionView {
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].Order < qs[j].Order
	})

	return qs
}

func QuestionsFromForm(formValues url.Values) []Question {
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

	return templatesQuestions
}

func QuestionsHaveErrors(qs []QuestionView) bool {
	for _, question := range qs {
		if question.Error != "" {
			return true
		}
	}

	return false
}

func ApplyAnsweringQuestionValidations(questionViews []QuestionView, formValues url.Values) []QuestionView {
	for i := range questionViews {
		val := formValues.Get(questionViews[i].ID.Hex())
		questionViews[i].Value = val

		if val == "" {
			questionViews[i].Error = "This field is required."
		}
	}

	return questionViews
}
