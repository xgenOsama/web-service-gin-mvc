package controllers

import (
	"net/http"

	"example.com/gin-mvc/initializers"
	"example.com/gin-mvc/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get the data of the request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// create a post
	post := models.Post{Title: body.Body, Body: body.Title}

	// insert into database
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostByID(c *gin.Context) {
	// get data from url
	id := c.Param("id")
	// get the post
	var post models.Post
	initializers.DB.First(&post, id)
	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get data from url
	id := c.Param("id")
	// get data from request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// get the post
	var post models.Post
	initializers.DB.First(&post, id)
	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDELETE(c *gin.Context) {
	// get data from url
	id := c.Param("id")
	// get the post
	var post models.Post
	initializers.DB.Delete(&post, id)
	// respond with it
	c.Status(200)
}
