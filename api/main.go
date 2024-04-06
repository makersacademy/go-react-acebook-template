package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/env"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
	"github.com/makersacademy/go-react-acebook-template/api/src/routes"
)

func main() {
	env.LoadEnv()

	app := setupApp()

	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	models.SeedUserIfNotExist()
	// Create a test testPost. Delete these lines when you are creating posts of your own.
	// testPost := models.Post{
	// 	Message: fmt.Sprintf("This is a test message created at %v!", time.Now()),
	// }
	// testPost.Save()

	app.Run(":8082")

}

func setupApp() *gin.Engine {
	app := gin.Default()
	setupCORS(app)
	routes.SetupRoutes(app)
	return app
}

func setupCORS(app *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}

	app.Use(cors.New(config))
}
