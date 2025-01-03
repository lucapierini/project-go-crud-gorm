package config

import (
	"github.com/lucapierini/project-go-crud-gorm/models"
)


func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}