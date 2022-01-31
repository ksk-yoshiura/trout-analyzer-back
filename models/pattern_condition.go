package models

import (
	"github.com/jinzhu/gorm"
)

type PatternCondition struct {
	gorm.Model
	TypeNum  int    `json:"type_num"`
	TypeName string `json:"type_name"`
}
