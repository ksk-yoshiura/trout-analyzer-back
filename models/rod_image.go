package models

import (
	"gorm.io/gorm"
)

type RodImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	RodId     uint   `json:"rod_id"`
}
