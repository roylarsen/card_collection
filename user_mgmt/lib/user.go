package lib

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

//UserAdd handles adding new users
func (client *Handler) UserAdd(c echo.Context) error {
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

//Login handles taking a user account and logging them in
func (client *Handler) Login(c echo.Context) error {
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
