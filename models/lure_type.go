package models

import (
	"github.com/jinzhu/gorm"
)

type LureType struct {
	gorm.Model
	TypeName string `json:"type_name"`
}
