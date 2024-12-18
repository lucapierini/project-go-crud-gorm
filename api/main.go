package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-crud-gorm/controllers"
	"github.com/lucapierini/project-go-crud-gorm/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", controllers.TestFunction)
	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()
}