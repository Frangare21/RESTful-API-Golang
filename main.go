package main

import (
	"restful-api/controllers"
	"restful-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreatefunc)
	router.GET("/posts", controllers.PostsIndexfunc)
	router.GET("/posts/:id", controllers.PostsByIdfunc)
	router.DELETE("/posts/:id", controllers.PostsDeletefunc)
	router.PUT("/posts/:id", controllers.PostsUpdatefunc)

	router.Run()
}
