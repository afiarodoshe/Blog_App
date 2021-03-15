package controllers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"main.go/config"
	"main.go/pkg/db/model"
	"net/http"
)

func AddPost(c echo.Context) error {
	post := &model.Post{
		Pid:    uuid.New().String(),
	}
	if err := c.Bind(post); err != nil {
		return err
	}
	fmt.Println(post)

	Postcollection := config.GetPostCollection()
	result, err := Postcollection.InsertOne(context.TODO(), post)
	log.Println(result.InsertedID)
	log.Println(err)
	if err != nil {
		return err
	}
	var returnMessage string
	returnMessage = "Post: " + post.PostTitle + " added successfully to database :)"
	return c.String(http.StatusCreated, returnMessage)
}

func GetPost(c echo.Context) error {
	postTitle := c.QueryParam("postTitle")
	var post bson.M
	Postcollection := config.GetPostCollection()
	Postcollection.FindOne(context.TODO(), bson.M{"postTitle": postTitle},
	).Decode(&post)
	if len(post) == 0{
		return c.String(http.StatusOK, "post Not Found! :(")
	}
	return c.JSON(http.StatusOK, post)
}

func UpdatePost(c echo.Context) error {
	post := &model.Post{}
	if err := c.Bind(post); err != nil {
		return err
	}

	postTitle := c.QueryParam("postTitle")
	Postcollection := config.GetPostCollection()

	var post1 bson.M
	Postcollection.FindOne(
		context.TODO(),
		bson.M{"postTitle": postTitle},
	).Decode(&post1)
	if len(post1) > 0 {
		result, err := Postcollection.UpdateOne(
			context.TODO(),
			bson.M{"postTitle": postTitle},
			bson.D{
				{"$set", bson.D{{"postBody", post.PostBody}}},
				{"$set", bson.D{{"postType", post.PostType}}},
			})
		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nUpdate Unsuccessful :("
			log.Fatal(err)
		} else if result.ModifiedCount == 0 {
			returnMessage = "post not found :("
		} else {
			returnMessage = post.PostTitle + " updated successfully :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	} else {
		return c.String(http.StatusOK, "Post Not Found! :(")
	}
}
