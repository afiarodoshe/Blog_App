package controllers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main.go/config"
	"main.go/pkg/db/model"
	"net/http"
	"regexp"
)

func GetUser(c echo.Context) error {
	email := c.QueryParam("email")
	var user bson.M
	usercollection := config.GetUserCollection()

	usercollection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&user)
	if user["userStatus"] == "D" {
		return c.String(http.StatusOK, "user Not Found! :(")
	}
	if len(user) > 0 {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.String(http.StatusOK, "user Not Found! :(")
	}
}


func GetAllUser(c echo.Context) error {
	email := c.QueryParam("email")
	Usercollection := config.GetUserCollection()

	opts := options.Find()
	opts.SetSort(bson.D{{"email", -1}})
	sortCursor, err := Usercollection.Find(context.TODO(), bson.D{{"email", bson.D{{"$gt", email}}}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var userSorted []bson.M
	if err = sortCursor.All(context.TODO(), &userSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(userSorted)
	return c.JSON(http.StatusOK, userSorted)
}



func AddUser(c echo.Context) error {
	user := &model.User{
		Uid:    uuid.New().String(),
		UserStatus: "V",
	}
	if err := c.Bind(user); err != nil {
		return err
	}

	fmt.Println(user)

	if !isEmailValid(user.Email) {
		return c.JSON(http.StatusOK, "Invalid Email :(")
	}

	Usercollection := config.GetUserCollection()

	var user1 bson.M
	Usercollection.FindOne(
		context.TODO(),
		bson.M{"email": user.Email},
	).Decode(&user1)
	fmt.Println(user1["email"])
	if len(user1) > 0 {
		return c.JSON(http.StatusOK, "Email address taken\n Use another :(")
	} else {
		result, err := Usercollection.InsertOne(context.TODO(), user)
		log.Println(result.InsertedID)
		log.Println(err)
		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nUser addition Unsuccessful :("
			log.Fatal(err)
		} else if result.InsertedID == 0 {
			returnMessage = "Something went wrong! \nUser addition Unsuccessful :("
		} else {
			returnMessage = "User: " + user.FullName + " added successfully to database :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	}
}

func UpdateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	email := c.QueryParam("email")
	Usercollection := config.GetUserCollection()

	var user1 bson.M
	Usercollection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&user1)
	if user1["status"] == "D" {
		return c.String(http.StatusOK, "User Not Found! :(")
	}
	if len(user1) > 0 {

		result, err := Usercollection.UpdateOne(
			context.TODO(),
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"fullName", user.FullName}}},
				{"$set", bson.D{{"username", user.UserName}}},
				{"$set", bson.D{{"userRole", user.UserRole}}},
			})
		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nUpdate Unsuccessful :("
			log.Fatal(err)
		} else if result.ModifiedCount == 0 {
			returnMessage = "User not found :("
		} else {
			returnMessage = user.Email + " updated successfully :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	} else {
		return c.String(http.StatusOK, "User Not Found! :(")
	}
}

func DeleteUser(c echo.Context) error {
	email := c.QueryParam("email")

	Usercollection := config.GetUserCollection()
	var user bson.M
	Usercollection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&user)
	fmt.Println(user["name"])
	if len(user) > 0 {
		_, err := Usercollection.UpdateOne(
			context.TODO(),
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"userStatus", "D"}}},
			})

		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nDelete Unsuccessful :("
			log.Fatal(err)
		} else {
			returnMessage = email + " deleted from database successfully :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	} else {
		return c.String(http.StatusOK, "User Not Found! :(")
	}
}


var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

