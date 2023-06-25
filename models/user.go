package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  *string            `bson:"username" json:"username" validate:"required,min=2,max=40"`
	Bio       *string            `json:"bio"`
	CreatedAt time.Time          `bson:"created_At" json:"created_At"`
	UpdatedAt time.Time          `bson:"updated_At" json:"updated_At"`
	UserId    string             `bson:"user_Id" json:"user_Id"`
}
