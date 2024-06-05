package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionType string

const (
	TextQuestion     QuestionType = "text"
	NumberQuestion   QuestionType = "number"
	SelectQuestion   QuestionType = "select"
	RangeQuestion    QuestionType = "range"
	TextAreaQuestion QuestionType = "text_area"
	RadioQuestion    QuestionType = "radio"
	CheckboxQuestion QuestionType = "checkbox"
)

type Question struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Label   string             `bson:"label"`
	Type    QuestionType       `bson:"type"`
	Options []string           `bson:"options,omitempty"`
	Min     int                `bson:"min,omitempty"`
	Max     int                `bson:"max,omitempty"`
	Order   int                `bson:"order"`
}

type Template struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	Username    string             `bson:"username"`
	Questions   []Question         `bson:"questions,omitempty"`
}
