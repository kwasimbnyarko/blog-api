package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwasimbnyarko/blog-api/controllers"
)

func PostRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/post", controllers.CreatePost())
	incomingRoutes.GET("/post/:postId", controllers.ViewPost())
	incomingRoutes.GET("/post", controllers.ViewAllPost())
	incomingRoutes.GET("/post/user/:username", controllers.ViewAllPostsFromUser())
	incomingRoutes.PUT("/post/:postId", controllers.UpdatePost())
	incomingRoutes.DELETE("/post/:postId", controllers.DeletePost())

}
