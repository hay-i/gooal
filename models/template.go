package models

import (
	"time"
)

type Template struct {
	Title     string    `bson:"title"`
	CreatedAt time.Time `bson:"created_at"`
}

type DefaultTemplate struct {
    ID        string    `bson:"_id"`
    // TODO: is there a way to embed the above in here?
	Title     string    `bson:"title"`
	CreatedAt time.Time `bson:"created_at"`
}
