package models

import (
	"gorm.io/gorm"
)

type FieldImage struct {
	gorm.Model
	ImageFile string `json:"image_file"`
	FieldId   uint   `json:"field_id"`
}
