package controllers

import (
	"github.com/danutryas/go-crud/initializers"
	model "github.com/danutryas/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a Post
	post := model.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	// Return

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsIndex(c *gin.Context) {
	// Get the post
	var posts []model.Post
	initializers.DB.Find(&posts)

	// Respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
func PostsShow(c *gin.Context) {
	// get id
	id := c.Param("id")

	// Get the post
	var post model.Post
	initializers.DB.First(&post, id)

	// Respond
	c.JSON(200, gin.H{
		"posts": post,
	})
}
func PostsUpdate(c *gin.Context) {
	// get id from params url
	id := c.Param("id")
	// get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// find post that will be update
	var post model.Post
	initializers.DB.First(&post, id)
	// update the post
	initializers.DB.Model(&post).Updates(model.Post{
		Body: body.Body, Title: body.Title,
	})

	// respond
}

func PostsDelete(c *gin.Context) {
	// get id from url params
	id := c.Param("id")
	// delete the post
	initializers.DB.Delete(&model.Post{}, id)

	// respond
	c.Status(200)

}
