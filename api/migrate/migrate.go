package main

import (
	"github.com/lucapierini/project-go-crud-gorm/models"
	"github.com/lucapierini/project-go-crud-gorm/initializers"
)
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}