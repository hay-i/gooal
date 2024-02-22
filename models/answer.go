package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	TemplateID primitive.ObjectID `bson:"template_id"`
	QuestionID primitive.ObjectID `bson:"question_id"`
	Answer     string             `bson:"answer"`
}
