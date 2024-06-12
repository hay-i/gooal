package views

import "github.com/hay-i/gooal/internal/models"

func QuestionAnswersFromForm(questionViews []QuestionView) []models.QuestionAnswer {
	questionAnswers := make([]models.QuestionAnswer, len(questionViews))

	for i, q := range questionViews {
		questionAnswers[i] = models.QuestionAnswer{
			QuestionID: q.ID,
			Answer:     q.Value,
		}
	}

	return questionAnswers
}
