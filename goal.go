package main

import "time"

type Goal struct {
	Title     string    `bson:"title"`
	CreatedAt time.Time `bson:"created_at"`
}
