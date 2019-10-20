package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/add", testFun)
	e.GET("/retr", retrFun)
	e.Logger.Fatal(e.Start(":1323"))
}

func testFun(c echo.Context) error {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	collection := client.Database("mydb").Collection("persons")

	roy := Person{"Roy", 33}

	insertResult, err := collection.InsertOne(context.TODO(), roy)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, insertResult.InsertedID)
}

func retrFun(c echo.Context) error {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	collection := client.Database("mydb").Collection("persons")
	filter := bson.D{{"name", "Roy"}}

	var result Person

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "Found a single document: "+result.Name+"\n")
}
