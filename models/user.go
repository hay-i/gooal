package models

// User represents the user model.
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"` // Hashed password
}
