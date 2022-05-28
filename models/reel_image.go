package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type ReelImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	ReelId    uint   `json:"reel_id"`
}

/**
  リール画像登録
*/
func CreateReelImage(image ReelImage) error {
	db := database.GetDBConn()
	result := db.Create(&image).Error
	return result
}
