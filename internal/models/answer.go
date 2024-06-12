package models

import (
	"time"

	"github.com/hay-i/gooal/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Answer struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	TemplateID      primitive.ObjectID `bson:"template_id"`
	Username        string             `bson:"username"`
	CreatedAt       time.Time          `bson:"created_at"`
	QuestionAnswers []QuestionAnswer   `bson:"questions,omitempty"`
}

func (a Answer) FromForm(templateID primitive.ObjectID, username string, questionAnswers []QuestionAnswer) Answer {
	a.TemplateID = templateID
	a.Username = username
	a.CreatedAt = time.Now()
	a.QuestionAnswers = questionAnswers

	return a
}

func (a Answer) Save(database *mongo.Database) {
	db.Save(database, "answers", a)
}
