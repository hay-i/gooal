package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionAnswer struct {
	QuestionID primitive.ObjectID `bson:"question_id"`
	Answer     string             `bson:"answer"`
}

type Answer struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	TemplateID      primitive.ObjectID `bson:"template_id"`
	Username        string             `bson:"username"`
	CreatedAt       time.Time          `bson:"created_at"`
	QuestionAnswers []QuestionAnswer   `bson:"questions,omitempty"`
}
