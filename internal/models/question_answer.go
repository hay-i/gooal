package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionAnswer struct {
	QuestionID primitive.ObjectID `bson:"question_id"`
	Answer     string             `bson:"answer"`
}

func QuestionAnswersFromForm(questionViews []QuestionView) []QuestionAnswer {
	questionAnswers := make([]QuestionAnswer, len(questionViews))

	for i, q := range questionViews {
		questionAnswers[i] = QuestionAnswer{
			QuestionID: q.ID,
			Answer:     q.Value,
		}
	}

	return questionAnswers
}
