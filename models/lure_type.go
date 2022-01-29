package models

import (
	"github.com/jinzhu/gorm"
)

type LureType struct {
	gorm.Model
	TypeName int `json:"type_name"`
}
