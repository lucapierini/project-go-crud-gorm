package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-crud-gorm/initializers"
	"github.com/lucapierini/project-go-crud-gorm/models"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"post": post,
		})
	}
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Post no encontrado",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post.Title = body.Title
	post.Body = body.Body
	
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "Post eliminado",
	})
}
