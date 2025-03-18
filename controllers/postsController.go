package controllers

import (
	"restful-api/initializers"
	"restful-api/models"

	"github.com/gin-gonic/gin"
)

func PostsCreatefunc(c *gin.Context) {
	//Get the post data from the request
	var req struct {
		Body  string
		Title string
	}

	c.Bind(&req)

	post := models.Post{Title: req.Title, Body: req.Body}

	result := initializers.DB.Create(&post)

	if result.Error == nil {
		c.JSON(200, gin.H{
			"message": "Post created succesfully!",
			"post":    post,
			"result":  result.RowsAffected,
		})
	} else {
		c.JSON(500, gin.H{
			"message": "Post not created!",
			"error":   result.Error,
		})
	}
}

func PostsIndexfunc(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsByIdfunc(c *gin.Context) {
	var post models.Post

	initializers.DB.First(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDeletefunc(c *gin.Context) {
	var post models.Post

	initializers.DB.Delete(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"message": "Post deleted successfully!",
	})
}

func PostsUpdatefunc(c *gin.Context) {
	var post models.Post

	initializers.DB.First(&post, c.Param("id"))

	var req struct {
		Body  string
		Title string
	}

	c.Bind(&req)

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Body != "" {
		post.Body = req.Body
	}

	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"message": "Post updated successfully!",
		"post":    post,
	})
}
