package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	Default     bool               `bson:"default,omitempty"`
}
