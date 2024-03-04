package db

import (
	"context"
	"fmt"
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
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 1", Description: ":)", Type: models.TextQuestion, Order: 1},
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 2", Description: ":)", Type: models.NumberQuestion, Order: 2},
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 3", Description: ":)", Type: models.SelectQuestion, Order: 3, Options: []string{"Option 1", "Option 2", "Option 3"}},
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 4", Description: ":)", Type: models.RangeQuestion, Order: 4, Min: 1, Max: 10},
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 5", Description: ":)", Type: models.TextAreaQuestion, Order: 5},
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 6", Description: ":)", Type: models.RadioQuestion, Order: 6, Options: []string{"Option 1", "Option 2", "Option 3"}},
				{ID: primitive.NewObjectID(), Title: "Template 1, Question 7", Description: ":)", Type: models.CheckboxQuestion, Order: 7, Options: []string{"Option 1", "Option 2", "Option 3"}},
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
				Answer:     fmt.Sprintf("Answer for: %s", question.Title),
			}

			filter := bson.M{"answer": answer.Answer}
			count, err := answerCollection.CountDocuments(ctx, filter)
			if err != nil {
				panic(err)
			}

			if count > 0 {
				continue
			}

			_, err = answerCollection.InsertOne(ctx, answer)
			if err != nil {
				panic(err)
			}
		}
	}
}
