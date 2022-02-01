package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type LureType struct {
	gorm.Model
	TypeName string `json:"type_name"`
}

/**
  ルアータイプ一覧取得
*/
func GetAllLureTypes(lure_types []LureType) []LureType {
	db := database.GetDBConn()
	db.Find(&lure_types)
	return lure_types
}
