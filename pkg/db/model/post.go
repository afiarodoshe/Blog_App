package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Pid   string `json:"pId" bson:"pId"`
	UserID        primitive.ObjectID `json:"U_id,omitempty" bson:"U_id,omitempty"`
	PostID        primitive.ObjectID `json:"P_id,omitempty" bson:"P_id,omitempty"`
	PostTitle string `json:"postTitle" bson:"postTitle"`
	PostBody  string `json:"postBody" bson:"postBody"`
	PostType  string `json:"postType" bson:"postType"`
	LikeCount int    `json:"likeCount" bson:"likeCount"`
}
