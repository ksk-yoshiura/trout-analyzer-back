package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type FishingLine struct {
	gorm.Model
	LineCondition ToolCondition `gorm:"foreignKey:LineTypeId"`
	Name          string        `json:"name"`
	UserId        int           `json:"userId"`
	LineTypeId    string        `json:"lineTypeId"`
	Thickness     int           `json:"thickness"`
	CompanyName   string        `json:"companyName"`
}

func FishingLineValidate(fishing_line FishingLine) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(fishing_line)
}

/**
  ライン一覧取得
*/
func GetAllLines(fishing_lines []FishingLine, uid int) []FishingLine {
	db := database.GetDBConn()
	// ログインユーザは自分のラインしか見れない
	db.Where("user_id = ?", uid).Preload("LineCondition").Find(&fishing_lines)
	return fishing_lines
}

/**
  ライン取得
*/
func GetLine(fishing_line FishingLine, line_id int, uid int) FishingLine {
	db := database.GetDBConn()
	// ログインユーザは自分のラインしか見れない
	db.Where("user_id = ?", uid).Preload("LineCondition").First(&fishing_line, line_id)
	return fishing_line
}

/**
  ライン更新
*/
func UpdateLine(f FishingLine, line_id int) error {
	var fishing_line FishingLine
	db := database.GetDBConn()

	db.First(&fishing_line, line_id)

	result := db.Model(&fishing_line).Updates(FishingLine{
		Name:        f.Name,
		UserId:      f.UserId,
		LineTypeId:  f.LineTypeId,
		Thickness:   f.Thickness,
		CompanyName: f.CompanyName,
	}).Error
	return result
}

/**
  ライン作成
*/
func CreateLine(reel FishingLine) error {
	db := database.GetDBConn()
	result := db.Create(&reel).Error
	return result
}

/**
  ライン削除
*/
func DeleteLine(fishing_line FishingLine, line_id int) error {
	db := database.GetDBConn()
	db.First(&fishing_line, line_id)
	result := db.Delete(&fishing_line).Error
	return result
}
