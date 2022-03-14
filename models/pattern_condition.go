package models

import (
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
)

type PatternCondition struct {
	gorm.Model
	TypeNum  int    `json:"typeNum"`
	TypeName string `json:"typeName"`
}

const (
	RESULT_TYPE_NUM  = 1
	LURE_SPEED_NUM   = 2
	LURE_DEPTH_NUM   = 3
	WEATHER_TYPE_NUM = 4
)

/**
  ヒットパターン条件各種一覧取得
*/
func GetAllPatternConditions(pattern_conditions []PatternCondition) []PatternCondition {
	db := database.GetDBConn()
	db.Find(&pattern_conditions)
	return pattern_conditions
}

/**
  釣果条件各種一覧取得
*/
func GetResultConditions(pattern_conditions []PatternCondition) []PatternCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", RESULT_TYPE_NUM).Find(&pattern_conditions)
	return pattern_conditions
}

/**
  ルアー速度条件各種一覧取得
*/
func GetLureSpeedConditions(pattern_conditions []PatternCondition) []PatternCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", LURE_SPEED_NUM).Find(&pattern_conditions)
	return pattern_conditions
}

/**
  ルアー深度条件各種一覧取得
*/
func GetLureDepthConditions(pattern_conditions []PatternCondition) []PatternCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", LURE_DEPTH_NUM).Find(&pattern_conditions)
	return pattern_conditions
}

/**
  天気条件各種一覧取得
*/
func GetWeatherConditions(pattern_conditions []PatternCondition) []PatternCondition {
	db := database.GetDBConn()
	db.Where("type_num = ?", WEATHER_TYPE_NUM).Find(&pattern_conditions)
	return pattern_conditions
}
