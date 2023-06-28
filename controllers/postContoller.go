package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kwasimbnyarko/blog-api/database"
	"github.com/kwasimbnyarko/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()

var postCollection = database.OpenCollection(database.Client, "post")

func Awake() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusCreated, "Server running \nVisit https://github.com/kwasimbnyarko/blog-api for details")
	}
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post

		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, bson.M{"error": "incorrect request body"})
		}

		post.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.Id = primitive.NewObjectID()
		post.PostId = post.Id.Hex()

		if validationErr := validate.Struct(post); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		_, err := postCollection.InsertOne(ctx, post)

		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "post creation failed"})
			log.Fatal(err)
		}
		c.JSON(http.StatusCreated, bson.M{"sucess": "post created"})
	}
}

func ViewPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
		postId := c.Param("postId")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		err := postCollection.FindOne(ctx, bson.M{"post_id": postId}).Decode(&post)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "error occured while finidng post"})
		}
		c.JSON(http.StatusOK, post)
	}
}

func ViewAllPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []models.Post
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		cursor, err := postCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "error occurred while getting employees"})
			log.Panic(err)
		}

		defer cancel()
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				log.Panic(err.Error())
			}
		}(cursor, ctx)

		for cursor.Next(ctx) {
			var post models.Post

			if err = cursor.Decode(&post); err != nil {
				log.Fatal(err)
			}
			posts = append(posts, post)
		}

		c.JSON(http.StatusOK, posts)
	}
}

func ViewAllPostsFromUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []models.Post
		username := c.Param("username")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		cursor, err := postCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "error occurred while getting employees"})
			log.Panic(err)
		}

		defer cancel()

		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				log.Panic(err.Error())
			}
		}(cursor, ctx)

		for cursor.Next(ctx) {
			var post models.Post

			if err = cursor.Decode(&post); err != nil {
				log.Fatal(err)
			}
			if post.Username == username {
				posts = append(posts, post)
			}
		}

		if posts == nil {
			c.JSON(http.StatusNotFound, bson.M{"msg": "user not found"})
			return
		}

		c.JSON(http.StatusOK, posts)
	}
}

func UpdatePost() gin.HandlerFunc {
	type body struct {
		Username *string `json:"username" bson:"username" validate:"required"`
		Title    *string `bson:"title" json:"title" validate:"required"`
		Text     *string `json:"text" bson:"text" validate:"required"`
	}
	return func(c *gin.Context) {

		postId := c.Param("postId")

		var post models.Post
		var body body

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, bson.M{"error": "incorrect request body"})
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		err := postCollection.FindOne(ctx, bson.M{"post_id": postId}).Decode(&post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "error occurred in finding post"})
			log.Fatal(err)
		}

		var updateObj primitive.D

		if *body.Username != "" {
			updateObj = append(updateObj, bson.E{Key: "username", Value: body.Username})
		}

		if *body.Title != "" {
			updateObj = append(updateObj, bson.E{Key: "title", Value: body.Title})
		}

		if *body.Text != "" {
			updateObj = append(updateObj, bson.E{Key: "text", Value: body.Text})
		}

		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_At", Value: updatedAt})

		upsert := true
		filter := bson.M{"post_id": postId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := postCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{Key: "$set", Value: updateObj}},
			&opt)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "update unsuccessfull"})

			log.Panicln(err)

		}

		c.JSON(http.StatusOK, result)
	}
}

func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("postId")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		result, err := postCollection.DeleteOne(ctx, bson.M{"post_id": postId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"error": "error occurred upon deletion"})
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, result)
	}
}
