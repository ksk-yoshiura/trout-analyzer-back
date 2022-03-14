package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type ToolCondition struct {
	gorm.Model
	TypeNum  int    `json:"typeNum"`
	TypeName string `json:"typeName"`
}

const (
	ROD_TYPE_NUM  = 1
	REEL_GEAR_NUM = 2
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
func GetRodHardnessConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", ROD_TYPE_NUM).Find(&tool_conditions)
	return tool_conditions
}

/**
  リールのギア条件各種一覧取得
*/
func GetReelGearConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", REEL_GEAR_NUM).Find(&tool_conditions)
	return tool_conditions
}

/**
  リールの型番条件各種一覧取得
*/
func GetReelTypeConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", REEL_TYPE_NUM).Find(&tool_conditions)
	return tool_conditions
}

/**
  ラインタイプ条件各種一覧取得
*/
func GetLineTypeConditions(tool_conditions []ToolCondition) []ToolCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", LINE_TYPE_NUM).Find(&tool_conditions)
	return tool_conditions
}
