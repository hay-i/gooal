package models

type User struct {
	Username string `bson:"username" form:"username"`
	Password string `bson:"password" form:"password"`
	// TODO: Figure out how to add a created at field
	// - it currently breaks login when binding credentials
}
