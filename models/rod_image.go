package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type RodImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	RodId     uint   `json:"rod_id"`
}

/**
  ロッド画像登録
*/
func CreateRodImagee(image RodImage) error {
	db := database.GetDBConn()
	result := db.Create(&image).Error
	return result
}
