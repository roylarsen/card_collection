package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/roylarsen/card_collection/user_mgmt/lib"
)

func main() {
	//Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(lib.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/user_add" {
				return true
			}
			return false
		},
	}))

	//DB initialization
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	clientObj, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		e.Logger.Fatal(http.StatusInternalServerError, err)
	}

	// Check the connection
	err = clientObj.Ping(context.TODO(), nil)

	if err != nil {
		e.Logger.Fatal(http.StatusInternalServerError, err)
	}

	client := &lib.Handler{DB: clientObj}

	// Initializing API endpoints
	e.POST("/user_add", client.UserAdd)
	e.POST("/login", client.Login)
	e.POST("/add_card", client.AddCard)
	e.Logger.Fatal(e.Start(":1323"))
}
