package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/controllers"
	"github.com/makersacademy/go-react-acebook-template/api/src/middleware"
)

func setupPostRoutes(baseRouter *gin.RouterGroup) {
	posts := baseRouter.Group("/posts")
	// This sets ip a new router group namd posts (with the baseRouter)
	// All routes defined within the 'posts.' group will have the prefix '/posts'

	posts.POST("", middleware.AuthenticationMiddleware, controllers.CreatePost)
	posts.GET("", middleware.AuthenticationMiddleware, controllers.GetAllPosts)
	posts.GET("/:id", middleware.AuthenticationMiddleware, controllers.GetSpecificPost)
	posts.PUT("/:id/likes", middleware.AuthenticationMiddleware, controllers.UpdatePostLikes)
	posts.DELETE("/:id/delete", middleware.AuthenticationMiddleware, controllers.DeletePost)
	//comments routes
	posts.GET("/:id/comments", middleware.AuthenticationMiddleware, controllers.GetAllCommentsByPostId)
	posts.POST("/:id/comments", middleware.AuthenticationMiddleware, controllers.CreateComment)
	posts.GET("/:id/comments/:comment_id", middleware.AuthenticationMiddleware, controllers.GetSpecificComment)
	posts.DELETE("/:id/comments/:comment_id/delete", middleware.AuthenticationMiddleware, controllers.DeleteComment)

}
