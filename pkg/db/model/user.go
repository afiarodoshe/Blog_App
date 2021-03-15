package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Uid        string             `json:"uId" bson:"uId"`
	FullName   string             `json:"fullName" bson:"fullName"`
	UserName   string             `json:"username" bson:"username"`
	Email      string             `json:"email" bson:"email"`
	UserStatus string             `json:"userStatus" bson:"userStatus"`
	UserRole   string             `json:"userRole" bson:"userRole"`
}
