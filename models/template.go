package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionType string

const (
	TextQuestion   QuestionType = "text"
	NumberQuestion QuestionType = "number"
	SelectQuestion QuestionType = "select"
)

type Question struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description,omitempty"`
	Type        QuestionType       `bson:"type"`
}

type Template struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	Default     bool               `bson:"default,omitempty"`
	Questions   []Question         `bson:"questions,omitempty"`
}
