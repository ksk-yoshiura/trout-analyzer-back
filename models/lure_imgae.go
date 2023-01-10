package models

import (
	"gorm.io/gorm"
)

type LureImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	LureId    uint   `json:"lure_id"`
}
