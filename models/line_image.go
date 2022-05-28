package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type LineImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	LineId    uint   `json:"line_id"`
}

/**
  ライン画像登録
*/
func CreateLineImage(image LineImage) error {
	db := database.GetDBConn()
	result := db.Create(&image).Error
	return result
}
