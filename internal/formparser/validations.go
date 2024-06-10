package formparser

import (
	"net/url"

	"github.com/hay-i/gooal/internal/models"
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
