package formparser

import (
	"net/url"
	"sort"

	"github.com/hay-i/gooal/models"
	"github.com/labstack/echo/v4"
)

func ValidateFormValues(c echo.Context) (url.Values, error) {
	if err := c.Request().ParseForm(); err != nil {
		return nil, err
	}

	formValues, err := c.FormParams()
	if err != nil {
		return nil, err
	}

	return formValues, nil
}

func QuestionsToView(qs []models.Question) []models.QuestionView {
	questionViews := make([]models.QuestionView, len(qs))
	for i, question := range qs {
		questionView := models.QuestionView{Question: question}
		questionViews[i] = questionView
	}

	// TODO: move to model?
	return sortQuestionsByOrder(questionViews)
}

func sortQuestionsByOrder(qs []models.QuestionView) []models.QuestionView {
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].Order < qs[j].Order
	})

	return qs
}

func ApplyValidations(qs []models.QuestionView, formValues url.Values) []models.QuestionView {
	for i := range qs {
		val := formValues.Get(qs[i].ID.Hex())
		qs[i].Value = val

		if val == "" {
			qs[i].Error = "This field is required."
		}
	}

	return qs
}

func HasErrors(qs []models.QuestionView) bool {
	for _, question := range qs {
		if question.Error != "" {
			return true
		}
	}

	return false
}
