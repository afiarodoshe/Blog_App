package controllers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"log"
	"main.go/config"
	"main.go/pkg/db/model"
	"net/http"
)

func AddComment(c echo.Context) error {
	comment := &model.Comment{
		Cid: uuid.New().String(),
	}
	if err := c.Bind(comment); err != nil {
		return err
	}
	fmt.Println(comment)

	Commentcollection := config.GetCommentCollection()
	result, err := Commentcollection.InsertOne(context.TODO(), comment)
	log.Println(result.InsertedID)
	log.Println(err)
	if err != nil {
		return err
	}
	var returnMessage string
	returnMessage = "Comment: " + comment.CommentDetails + " added successfully to database :)"
	return c.String(http.StatusCreated, returnMessage)
}

