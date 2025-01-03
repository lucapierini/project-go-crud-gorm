package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-crud-gorm/config"
	"github.com/lucapierini/project-go-crud-gorm/controllers"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	config.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", controllers.TestFunction)
	posts := r.Group("/posts")
	{
		posts.GET("/", controllers.PostsIndex)
		posts.GET("/:id", controllers.PostsShow)
		posts.PUT("/:id", controllers.PostUpdate)
		posts.DELETE("/:id", controllers.PostDelete)
		posts.POST("/", controllers.PostCreate)
	}

	r.Run()
}
