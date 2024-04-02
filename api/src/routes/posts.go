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
}

// func setupLikePostRoute(baseRouter *gin.RouterGroup) {
// 	post := baseRouter.Group("/posts/:id/likes")

// }
