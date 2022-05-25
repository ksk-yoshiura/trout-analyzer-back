package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type FieldImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	FieldId   uint   `json:"field_id"`
}

/**
  フィールド作成
*/
func CreateFieldImage(image FieldImage) error {
	db := database.GetDBConn()
	result := db.Create(&image).Error
	return result
}
