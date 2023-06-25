package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, bson.M{"success": "hehe"})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
