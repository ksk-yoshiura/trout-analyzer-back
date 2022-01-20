package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/wcl48/valval"
)

type Rod struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	CompanyName string `json:"company_name"`
	Length      string `json:"length"`
	HardnessId  int    `json:"hardness_id"`
}

func RodValidate(rod Rod) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(rod)
}

/**
  ロッド一覧取得
*/
func GetAllRods(rods []Rod) []Rod {
	db := database.GetDBConn()
	db.Find(&rods)
	return rods
}

/**
  ロッド取得
*/
func GetRod(rod Rod, rod_id int) Rod {
	db := database.GetDBConn()
	db.First(&rod, rod_id)
	return rod
}

/**
  ロッド更新
*/
func UpdateRod(rod Rod, rod_id int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&rod, rod_id)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	hardness_id, _ := strconv.Atoi(c.FormValue("hardness_id"))
	length := c.FormValue("length")
	company_name := c.FormValue("company_name")

	result := db.Model(&rod).Updates(Rod{
		Name:        name,
		UserId:      user_id,
		HardnessId:  hardness_id,
		Length:      length,
		CompanyName: company_name,
	}).Error
	return result
}

/**
  ロッド作成
*/
func CreateRod(rod Rod) error {
	db := database.GetDBConn()
	result := db.Create(&rod).Error
	return result
}

/**
  ロッド削除
*/
func DeleteRod(rod Rod, rod_id int) error {
	db := database.GetDBConn()
	db.First(&rod, rod_id)
	result := db.Delete(&rod).Error
	return result
}
