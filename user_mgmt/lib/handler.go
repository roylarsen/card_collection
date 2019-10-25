package lib

import "go.mongodb.org/mongo-driver/mongo"

//Account Struct holds User account information
type Account struct {
	Name         string
	PasswordHash string
}

// Handler struct is a handler to pass DB client to functions
type (
	Handler struct {
		DB *mongo.Client
	}
)
