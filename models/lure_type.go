package models

import (
	"trout-analyzer-back/database"

	"gorm.io/gorm"
)

type LureType struct {
	gorm.Model
	TypeName string `json:"typeName"`
}

/**
  ルアータイプ一覧取得
*/
func GetAllLureTypes(lure_types []LureType) []LureType {
	db := database.GetDBConn()
	db.Find(&lure_types)
	return lure_types
}
