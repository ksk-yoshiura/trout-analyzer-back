package models

import (
	"gorm.io/gorm"
)

type ReelImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	ReelId    uint   `json:"reel_id"`
}
