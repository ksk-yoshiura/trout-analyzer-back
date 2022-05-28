package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
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
