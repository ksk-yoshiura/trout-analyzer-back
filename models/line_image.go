package models

import (
	"gorm.io/gorm"
)

type LineImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	LineId    uint   `json:"line_id"`
}
