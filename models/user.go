package models

// User represents the user model.
type User struct {
	Username string `bson:"username" form:"username"`
	Password string `bson:"password" form:"password"`
}
