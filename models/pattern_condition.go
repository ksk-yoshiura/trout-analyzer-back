package models

import (
	"github.com/jinzhu/gorm"
)

type PatternCondition struct {
	gorm.Model
	TypeNum  int `json:"type_num"`
	TypeName int `json:"type_name"`
}
