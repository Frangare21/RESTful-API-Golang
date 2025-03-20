package main

import (
	"restful-api/controllers"
	"restful-api/initializers"
	middlewares "restful-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.GenerateToken)
		api.POST("/register", controllers.RegisterUser)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			// Posts Routes
			secured.POST("/posts", controllers.PostsCreatefunc)
			secured.GET("/posts", controllers.PostsIndexfunc)
			secured.GET("/posts/:id", controllers.PostsByIdfunc)
			secured.DELETE("/posts/:id", controllers.PostsDeletefunc)
			secured.PUT("/posts/:id", controllers.PostsUpdatefunc)
		}
	}

	return router
}

func main() {

	router := initRouter()
	router.Run()

}
