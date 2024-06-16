package views

import (
	"sort"

	"github.com/hay-i/gooal/internal/models"
)

type QuestionView struct {
	models.Question
	Error string
	Value string
}

func SortQuestionsByOrder(qs []QuestionView) []QuestionView {
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].Order < qs[j].Order
	})

	return qs
}

func QuestionsHaveErrors(qs []QuestionView) bool {
	for _, question := range qs {
		if question.Error != "" {
			return true
		}
	}

	return false
}

func QuestionsToView(qs []models.Question) []QuestionView {
	questionViews := make([]QuestionView, len(qs))
	for i, question := range qs {
		questionView := QuestionView{Question: question}
		questionViews[i] = questionView
	}

	return SortQuestionsByOrder(questionViews)
}
