package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type HitPattern struct {
	gorm.Model
	UserId   int `json:"user_id"`
	LureId   int `json:"lure_id"`
	TackleId int `json:"tackle_id"`
	Speed    int `json:"speed"`
	Depth    int `json:"depth"`
	Weather  int `json:"weather"`
	Result   int `json:"result"`
	FieldId  int `json:"field_id"`
}

func HitPatternValidate(hit_pattern HitPattern) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(hit_pattern)
}

/**
  ヒットパターン一覧取得
*/
func GetAllHitPatterns(hit_patterns []HitPattern, uid int) []HitPattern {
	db := database.GetDBConn()
	// ログインユーザは自分のヒットパターンしか見れない
	db.Where("user_id = ?", uid).Find(&hit_patterns)
	return hit_patterns
}

/**
  ヒットパターン取得
*/
func GetHitPattern(hit_pattern HitPattern, hit_pattern_id int, uid int) HitPattern {
	db := database.GetDBConn()
	// ログインユーザは自分のヒットパターンしか見れない
	db.Where("user_id = ?", uid).First(&hit_pattern, hit_pattern_id)
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
		FieldId:  h.FieldId,
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
