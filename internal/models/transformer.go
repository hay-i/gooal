package models

func QuestionsToView(qs []Question) []QuestionView {
	questionViews := make([]QuestionView, len(qs))
	for i, question := range qs {
		questionView := QuestionView{Question: question}
		questionViews[i] = questionView
	}

	return SortQuestionsByOrder(questionViews)
}
