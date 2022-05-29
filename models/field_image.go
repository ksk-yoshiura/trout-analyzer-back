package models

import (
	"trout-analyzer-back/database"

	"gorm.io/gorm"
)

type FieldImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	FieldId   uint   `json:"field_id"`
}

/**
フィールド画像登録
*/
func CreateFieldImage(image FieldImage) error {
	db := database.GetDBConn()
	result := db.Create(&image).Error
	return result
}
