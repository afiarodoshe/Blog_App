package main

type User struct {
	UserId   int    `json:"userId" bson:"userId"`
	FullName string `json:"fullName" bson:"fullName"`
	UserName string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Status   string `json:"status" bson:"status"`
	UserRole string `json:"userRole" bson:"userRole"`
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
