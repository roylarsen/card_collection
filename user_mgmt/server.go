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

//Account Struct holds User account information
type Account struct {
	Name         string
	PasswordHash string
}

// Client struct is a handler to pass DB client to functions
type (
	Handler struct {
		DB *mongo.Client
	}
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

	client := &Handler{DB: clientObj}

	// Initializing API endpoints
	e.GET("/user_add", client.userAdd)
	e.GET("/login", client.login)
	e.Logger.Fatal(e.Start(":1323"))
}

func (client *Handler) userAdd(c echo.Context) error {
	collection := client.DB.Database("mydb").Collection("accounts")

	err := client.DB.Ping(context.TODO(), nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No conn to DB")
	}

	roy := Account{"Roy", "not a hash"}

	insertResult, err := collection.InsertOne(context.TODO(), roy)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, insertResult.InsertedID)
}

func (client *Handler) login(c echo.Context) error {
	collection := client.DB.Database("mydb").Collection("accounts")
	filter := bson.D{{"name", "Roy"}}

	var result Account

	err := client.DB.Ping(context.TODO(), nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No conn to DB")
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "Found a single document: "+result.Name+"\n")
}
