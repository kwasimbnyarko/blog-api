package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwasimbnyarko/blog-api/controllers"
)

func PostRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/api/post", controllers.CreatePost())
	incomingRoutes.GET("/api/post/:postId", controllers.ViewPost())
	incomingRoutes.GET("/api/post", controllers.ViewAllPost())
	incomingRoutes.GET("/api/post/user/:username", controllers.ViewAllPostsFromUser())
	incomingRoutes.PUT("/api/post/:postId", controllers.UpdatePost())
	incomingRoutes.DELETE("/api/post/:postId", controllers.DeletePost())

}
