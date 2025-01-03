package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-crud-gorm/config"
	"github.com/lucapierini/project-go-crud-gorm/models"
	"github.com/lucapierini/project-go-crud-gorm/errors"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	if c.Bind(&body) != nil {
		errors.RespondWithError(c, errors.ErrBadRequest)
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := config.DB.Create(&post)

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
	config.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	config.DB.First(&post, id)

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
	config.DB.First(&post, id)

	var body struct {
		Body  string
		Title string
	}

	if c.Bind(&body) != nil {
		errors.RespondWithError(c, errors.ErrBadRequest)
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	config.DB.Model(&post).Updates(models.Post{
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
	config.DB.First(&post, id)

	config.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "Post eliminado",
	})
}
