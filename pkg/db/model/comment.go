package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Cid   string `json:"cId" bson:"cId"`
	UserID        primitive.ObjectID `json:"U_id,omitempty" bson:"U_id,omitempty"`
	CommentID        primitive.ObjectID `json:"C_id,omitempty" bson:"C_id,omitempty"`
	PostID        primitive.ObjectID `json:"P_id,omitempty" bson:"P_id,omitempty"`
	CommentDetails string `json:"commentDetails" bson:"commentDetails"`
}

