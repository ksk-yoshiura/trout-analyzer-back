package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/wcl48/valval"
)

type Lure struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	LureTypeId  int    `json:"lure_type_id"`
	CompanyName string `json:"company_name"`
	Weight      string `json:"weight"`
	Color       string `json:"color"`
}

func LureValidate(lure Lure) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(lure)
}

/**
  ルアー一覧取得
*/
func GetAllLures(lures []Lure) []Lure {
	db := database.GetDBConn()
	db.Find(&lures)
	return lures
}

/**
  ルアー取得
*/
func GetLure(lure Lure, lure_id int) Lure {
	db := database.GetDBConn()
	db.First(&lure, lure_id)
	return lure
}

/**
  ルアー更新
*/
func UpdateLure(lure Lure, lure_id int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&lure, lure_id)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	lure_type_id, _ := strconv.Atoi(c.FormValue("lure_type_id"))
	company_name := c.FormValue("company_name")
	color := c.FormValue("color")
	weight := c.FormValue("weight")

	result := db.Model(&lure).Updates(Lure{
		Name:        name,
		UserId:      user_id,
		LureTypeId:  lure_type_id,
		CompanyName: company_name,
		Color:       color,
		Weight:      weight,
	}).Error
	return result
}

/**
  ルアー作成
*/
func CreateLure(lure Lure) error {
	db := database.GetDBConn()
	result := db.Create(&lure).Error
	return result
}

/**
  ルアー削除
*/
func DeleteLure(lure Lure, lure_id int) error {
	db := database.GetDBConn()
	db.First(&lure, lure_id)
	result := db.Delete(&lure).Error
	return result
}
