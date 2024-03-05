package models

import "time"

type User struct {
	Username  string    `bson:"username" form:"username"`
	Password  string    `bson:"password" form:"password"`
	CreatedAt time.Time `bson:"created_at"`
}
