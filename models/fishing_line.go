package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/wcl48/valval"
)

type FishingLine struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	LineTypeId  int    `json:"line_type_id"`
	Thickness   int    `json:"thickness"`
	CompanyName string `json:"company_name"`
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
	db.Where("user_id = ?", uid).Find(&fishing_lines)
	return fishing_lines
}

/**
  ライン取得
*/
func GetLine(fishing_line FishingLine, line_id int) FishingLine {
	db := database.GetDBConn()
	db.First(&fishing_line, line_id)
	return fishing_line
}

/**
  ライン更新
*/
func UpdateLine(fishing_line FishingLine, line_id int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&fishing_line, line_id)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	line_type_id, _ := strconv.Atoi(c.FormValue("line_type_id"))
	thickness, _ := strconv.Atoi(c.FormValue("thickness"))
	company_name := c.FormValue("company_name")

	result := db.Model(&fishing_line).Updates(FishingLine{
		Name:        name,
		UserId:      user_id,
		LineTypeId:  line_type_id,
		Thickness:   thickness,
		CompanyName: company_name,
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
