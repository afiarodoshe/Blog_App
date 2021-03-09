package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"time"
)

type User struct {
	UserId     int    `json:"userId" bson:"userId"`
	FullName   string `json:"fullName" bson:"fullName"`
	UserName   string `json:"username" bson:"username"`
	Email      string `json:"email" bson:"email"`
	UserStatus string `json:"userStatus" bson:"userStatus"`
	UserRole   string `json:"userRole" bson:"userRole"`
}

type Post struct {
	PostId    int    `json:"postId" bson:"postId"`
	PostTitle string `json:"postTitle" bson:"posTitle"`
	PostBody  string `json:"postBody" bson:"postBody"`
	PostType  string `json:"postType" bson:"postType"`
	LikeCount int    `json:"likeCount" bson:"likeCount"`
	UserRole  string `json:"userRole" bson:"userRole"`
}

type Comment struct {
	CommentId      int    `json:"commentId" bson:"commentId"`
	CommentDetails string `json:"commentDetails" bson:"commentDetails"`
	CommentStatus  string `json:"commentStatus" bson:"commentStatus"`
}

var client *mongo.Client
var err error

func main() {

	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	fmt.Println(databases)

	router := mux.NewRouter()

	router.HandleFunc("/posts", addUser).Methods("POST")

	router.HandleFunc("/posts", addPost).Methods("POST")

	router.HandleFunc("/posts/{Id}", getPost).Methods("GET")

	router.HandleFunc("/posts", updatePost).Methods("PUT")

	router.HandleFunc("/posts/{Id}", deletePost).Methods("DELETE")

	router.HandleFunc("/posts", addComment).Methods("POST")

	http.ListenAndServe(":5000", router)

}

func addUser(w http.ResponseWriter, r *http.Request) {

}
func addPost(w http.ResponseWriter, r *http.Request) {

}
func getPost(w http.ResponseWriter, r *http.Request) {

}
func updatePost(w http.ResponseWriter, r *http.Request) {

}
func deletePost(w http.ResponseWriter, r *http.Request) {

}
func addComment(w http.ResponseWriter, r *http.Request) {

}
