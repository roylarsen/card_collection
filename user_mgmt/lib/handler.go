package lib

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Account Struct holds User account information
type Account struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"passwordhash,omitempty" bson:"passwordhash"`
	Token        string             `json:"token,omitempty" bson:"-"`
}

// Handler struct is a handler to pass DB client to functions
type (
	Handler struct {
		DB *mongo.Client
	}
)

const (
	// Key (shhhh).
	Key = "secret"
)
