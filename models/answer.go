package models

// TODO: Currently unused

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	TemplateID primitive.ObjectID `bson:"template_id"`
	QuestionID primitive.ObjectID `bson:"question_id"`
	Answer     string             `bson:"answer"`
}

func AnswersForQuestion(answers []Answer, questionID primitive.ObjectID) []Answer {
	var result []Answer

	for _, answer := range answers {
		if answer.QuestionID == questionID {
			result = append(result, answer)
		}
	}

	return result
}
