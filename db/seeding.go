package db

import (
	"context"
	"time"

	"github.com/hay-i/chronologger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func seedTemplates(ctx context.Context, database *mongo.Database) {
	templateCollection := database.Collection("templates")

	defaultTemplates := []models.Template{
		{
			Title:       "Default Template #1",
			Description: "My description 1",
			CreatedAt:   time.Now(),
			Default:     true,
			Questions: []models.Question{
				{Title: "Question 1", Description: "My description 1", Type: models.TextQuestion},
				{Title: "Question 2", Description: "My description 2", Type: models.NumberQuestion},
				{Title: "Question 3", Description: "My description 3", Type: models.SelectQuestion},
			},
		},
		{
			Title:       "Default Template #2",
			Description: "My description 2",
			CreatedAt:   time.Now(),
			Default:     true,
			Questions: []models.Question{
				{Title: "Question 1", Description: "My description 1", Type: models.TextQuestion},
				{Title: "Question 2", Description: "My description 2", Type: models.NumberQuestion},
				{Title: "Question 3", Description: "My description 3", Type: models.SelectQuestion},
			},
		},
		{
			Title:       "Default Template #3",
			Description: "My description 3",
			CreatedAt:   time.Now(),
			Default:     true,
			Questions: []models.Question{
				{ID: primitive.NewObjectID(), Title: "Question 1", Description: "My description 1", Type: models.TextQuestion},
				{ID: primitive.NewObjectID(), Title: "Question 2", Description: "My description 2", Type: models.NumberQuestion},
				{ID: primitive.NewObjectID(), Title: "Question 3", Description: "My description 3", Type: models.SelectQuestion},
			},
		},
	}

	for _, defaultTemplate := range defaultTemplates {
		filter := bson.M{"title": defaultTemplate.Title}
		count, err := templateCollection.CountDocuments(ctx, filter)
		if err != nil {
			panic(err)
		}

		if count > 0 {
			continue
		}

		_, err = templateCollection.InsertOne(ctx, defaultTemplate)
		if err != nil {
			panic(err)
		}
	}
}

func seedAnswers(ctx context.Context, database *mongo.Database) {
	templateCollection := database.Collection("templates")

	templates := []models.Template{}

	cursor, err := templateCollection.Find(ctx, bson.M{"default": true})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &templates); err != nil {
		panic(err)
	}

	answerCollection := database.Collection("answers")

	for _, template := range templates {
		for _, question := range template.Questions {
			answer := models.Answer{
				TemplateID: template.ID,
				QuestionID: question.ID,
				Answer:     "My answer",
			}
			_, err = answerCollection.InsertOne(ctx, answer)
			if err != nil {
				panic(err)
			}
		}
	}
}
