package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username" json:"username" validate:"required,min=2,max=40"`
	Text      *string            `bson:"text" json:"text" validate:"required,max=140"`
	CreatedAt time.Time          `bson:"created_At" json:"created_At"`
	UpdatedAt time.Time          `bson:"updated_At" json:"updated_At"`
	PostId    string             `bson:"post_id" json:"post_id"`
}
