package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client
var e error

func LoadDB() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))

	// Connect to MongoDB
	client, e = mongo.Connect(context.TODO(), clientOptions)
	CheckError(e)

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	CheckError(e)

	return client, e
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func LoadEnvironments()  {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetUserCollection() *mongo.Collection {
	client, _ := LoadDB()
	// get collection as ref
	Usercollection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USER_COLLECTION"))
	return Usercollection
}
func GetPostCollection() *mongo.Collection {
	client, _ := LoadDB()
	// get collection as ref
	Postcollection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("POST_COLLECTION"))
	return Postcollection
}
func GetCommentCollection() *mongo.Collection {
	client, _ := LoadDB()
	// get collection as ref
	Commentcollection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COMMENT_COLLECTION"))
	return Commentcollection
}