package models

import (
	"time"
)

type Template struct {
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
}

// TODO: is there a way to embed Template in here?
type DefaultTemplate struct {
	ID          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
}
