package models

import (
	"trout-analyzer-back/database"

	"gorm.io/gorm"
)

type HitPattern struct {
	gorm.Model
	Lure             Lure             `gorm:"foreignKey:LureId"`
	Tackle           Tackle           `gorm:"foreignKey:TackleId"`
	Record           Record           `gorm:"foreignKey:RecordId"`
	SpeedCondition   PatternCondition `gorm:"foreignKey:Speed"`
	DepthCondition   PatternCondition `gorm:"foreignKey:Depth"`
	WeatherCondition PatternCondition `gorm:"foreignKey:Weather"`
	ResultCondition  PatternCondition `gorm:"foreignKey:Result"`
	UserId           int              `json:"userId"`
	LureId           int              `json:"lureId"`
	TackleId         int              `json:"tackleId"`
	Speed            int              `json:"speed"`
	Depth            int              `json:"depth"`
	Weather          int              `json:"weather"`
	Result           int              `json:"result"`
	RecordId         int              `json:"recordId"`
}

/**
  ヒットパターン一覧取得
*/
func GetAllHitPatterns(hit_patterns []HitPattern, uid int, record_id int) []HitPattern {
	db := database.GetDBConn()
	// ログインユーザは自分のヒットパターンしか見れない
	db.Where("user_id = ? AND record_id = ?", uid, record_id).Preload("Lure.LureType").Preload("Lure.LureImage").Preload("SpeedCondition").Preload("DepthCondition").Preload("WeatherCondition").Preload("ResultCondition").Find(&hit_patterns)
	return hit_patterns
}

/**
  ヒットパターン取得
*/
func GetHitPattern(hit_pattern HitPattern, hit_pattern_id int, uid int) HitPattern {
	db := database.GetDBConn()
	// ログインユーザは自分のヒットパターンしか見れない
	db.Where("user_id = ?", uid).Preload("Lure.Color").Preload("Lure.LureType").Preload("Lure.LureImage").Preload("Tackle.Reel.GearCondition").Preload("Tackle.Reel.TypeNumberCondition").Preload("Tackle.Rod.RodHardnessCondition").Preload("Tackle.Line.LineCondition").Preload("Record.Field").Preload("SpeedCondition").Preload("DepthCondition").Preload("WeatherCondition").Preload("ResultCondition").First(&hit_pattern, hit_pattern_id)
	return hit_pattern
}

/**
  ヒットパターン更新
*/
func UpdateHitPattern(h HitPattern, hit_pattern_id int) error {
	var hit_pattern HitPattern
	db := database.GetDBConn()

	// ログインユーザは自分のロッドしか見れない
	db.Where("user_id = ?", h.UserId).First(&hit_pattern, hit_pattern_id)

	result := db.Model(&hit_pattern).Updates(HitPattern{
		UserId:   h.UserId,
		LureId:   h.LureId,
		TackleId: h.TackleId,
		Speed:    h.Speed,
		Depth:    h.Depth,
		Weather:  h.Weather,
		Result:   h.Result,
		RecordId: h.RecordId,
	}).Error

	return result
}

/**
  ヒットパターン作成
*/
func CreateHitPattern(hit_pattern HitPattern) error {
	db := database.GetDBConn()
	result := db.Create(&hit_pattern).Error
	return result
}

/**
  ヒットパターン削除
*/
func DeleteHitPattern(hit_pattern HitPattern, hit_pattern_id int) error {
	db := database.GetDBConn()
	db.First(&hit_pattern, hit_pattern_id)
	result := db.Delete(&hit_pattern).Error
	return result
}
