package lib

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserAdd handles adding new users
func (client *Handler) UserAdd(c echo.Context) error {
	//Unpack post data to a user struct
	a := &Account{ID: primitive.NewObjectID()}
	if err := c.Bind(a); err != nil {
		return err
	}

	//Verify that we're not passing in empty data
	//if a.Email == "" || a.PasswordHash == "" {
	//	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid name or password"}
	//}
	if a.Email == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid user name"}
	}

	if a.PasswordHash == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid pass hash"}
	}

	collection := client.DB.Database("mydb").Collection("accounts")

	err := client.DB.Ping(context.TODO(), nil)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "No conn to DB"}
	}

	insertResult, err := collection.InsertOne(context.TODO(), a)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, insertResult.InsertedID)
}

//Login handles taking a user account and logging them in
func (client *Handler) Login(c echo.Context) error {
	a := new(Account)
	if err := c.Bind(a); err != nil {
		return err
	}

	collection := client.DB.Database("mydb").Collection("accounts")
	filter := bson.M{"email": a.Email, "passwordhash": a.PasswordHash}

	var result Account

	err := client.DB.Ping(context.TODO(), nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No conn to DB")
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = a.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	a.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	a.PasswordHash = "redacted" // Don't send password
	return c.JSON(http.StatusOK, a)
}

// AddCard takes a JWT and card info and adds the card to your collection
func (client *Handler) AddCard(c echo.Context) error {
	postData := echo.Map{}

	if err := c.Bind(&postData); err != nil {
		return err
	}

	if _, ok := postData["Name"]; ok {
		cardData := makeLookupReq(string(postData["Name"].(string)))

		return c.JSON(http.StatusOK, cardData)
	} else if _, ok = postData["Set"]; ok {
		cardData := makeLookupReq(string(postData["Set"].(string)) + "/" + string(postData["Num"].(string)))

		return c.JSON(http.StatusOK, cardData)
	}
	return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "No proper card data"}
}

func userIDFromToken(c echo.Context) string {
	account := c.Get("user").(*jwt.Token)
	claims := account.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func makeLookupReq(lookupString string) *echo.Map {
	url := "http://lookup:1323/" + lookupString

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	cardData := echo.Map{}

	_ = json.Unmarshal([]byte(body), &cardData)

	return &cardData
}
