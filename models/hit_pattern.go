package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
func GetAllHitPatterns(hit_patterns []HitPattern) []HitPattern {
	db := database.GetDBConn()
	db.Find(&hit_patterns)
	return hit_patterns
}

/**
  ヒットパターン取得
*/
func GetHitPattern(hit_pattern HitPattern, hit_pattern_id int) HitPattern {
	db := database.GetDBConn()
	db.First(&hit_pattern, hit_pattern_id)
	return hit_pattern
}

/**
  ヒットパターン更新
*/
func UpdateHitPattern(hit_pattern HitPattern, hit_pattern_id int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&hit_pattern, hit_pattern_id)
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	lure_id, _ := strconv.Atoi(c.FormValue("lure_id"))
	tackle_id, _ := strconv.Atoi(c.FormValue("tackle_id"))
	speed, _ := strconv.Atoi(c.FormValue("speed"))
	depth, _ := strconv.Atoi(c.FormValue("depth"))
	weather, _ := strconv.Atoi(c.FormValue("weather"))
	result_id, _ := strconv.Atoi(c.FormValue("result"))
	field_id, _ := strconv.Atoi(c.FormValue("field_id"))

	result := db.Model(&hit_pattern).Updates(HitPattern{
		UserId:   user_id,
		LureId:   lure_id,
		TackleId: tackle_id,
		Speed:    speed,
		Depth:    depth,
		Weather:  weather,
		Result:   result_id,
		FieldId:  field_id,
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
