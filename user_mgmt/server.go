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
	e.GET("/user_add", client.UserAdd)
	e.GET("/login", client.Login)
	e.Logger.Fatal(e.Start(":1323"))
}
