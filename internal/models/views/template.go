package views

type TemplateView struct {
	Title            string
	TitleError       string
	Description      string
	DescriptionError string
	QuestionViews    []QuestionView
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
