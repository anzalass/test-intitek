package utils

import (
	"example.com/m/v2/models"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(models.ProductModel{}, models.UserModel{})
	if err != nil {
		return
	}
}