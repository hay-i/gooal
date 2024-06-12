package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionAnswer struct {
	QuestionID primitive.ObjectID `bson:"question_id"`
	Answer     string             `bson:"answer"`
}
