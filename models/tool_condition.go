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

const (
	ROD_TYPE_NUM  = 1
	GEAR_TYPE_NUM = 2
	REEL_TYPE_NUM = 3
	LINE_TYPE_NUM = 4
)

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
	db.Where("type_num = ?", ROD_TYPE_NUM).Find(&tool_conditions)
	return tool_conditions
}

/**
  リールのギア条件各種一覧取得
*/
func GetReelGearConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", GEAR_TYPE_NUM).Find(&tool_conditions)
	return tool_conditions
}
