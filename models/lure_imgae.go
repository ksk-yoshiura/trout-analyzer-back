package models

import (
	"trout-analyzer-back/database"

	"gorm.io/gorm"
)

type LureImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	LureId    uint   `json:"lure_id"`
}

/**
  ルアー画像登録
*/
func CreateLureImage(image LureImage) error {
	db := database.GetDBConn()
	result := db.Create(&image).Error
	return result
}
