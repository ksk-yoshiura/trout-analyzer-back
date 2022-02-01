package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type ToolCondition struct {
	gorm.Model
	TypeNum  int    `json:"type_num"`
	TypeName string `json:"type_name"`
}

/**
  ツール条件各種一覧取得
*/
func GetAllToolConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Find(&tool_conditions)
	return tool_conditions
}

/**
  ロッドの硬さ条件各種一覧取得
*/
func GetAllRodHardnessConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", 1).Find(&tool_conditions)
	return tool_conditions
}
