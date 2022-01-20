package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/wcl48/valval"
)

type Reel struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	CompanyName string `json:"company_name"`
	TypeNumber  int    `json:"type_number"`
	Gear        string `json:"gear"`
}

func ReelValidate(reel Reel) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(reel)
}

/**
  リール一覧取得
*/
func GetAllReels(reels []Reel) []Reel {
	db := database.GetDBConn()
	db.Find(&reels)
	return reels
}

/**
  リール取得
*/
func GetReel(reel Reel, reel_id int) Reel {
	db := database.GetDBConn()
	db.First(&reel, reel_id)
	return reel
}

/**
  リール更新
*/
func UpdateReel(reel Reel, reel_id int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&reel, reel_id)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	type_number, _ := strconv.Atoi(c.FormValue("type_number"))
	gear := c.FormValue("gear")
	company_name := c.FormValue("company_name")

	result := db.Model(&reel).Updates(Reel{
		Name:        name,
		UserId:      user_id,
		TypeNumber:  type_number,
		Gear:        gear,
		CompanyName: company_name,
	}).Error
	return result
}

/**
  リール作成
*/
func CreateReel(reel Reel) error {
	db := database.GetDBConn()
	result := db.Create(&reel).Error
	return result
}

/**
  リール削除
*/
func DeleteReel(reel Reel, reel_id int) error {
	db := database.GetDBConn()
	db.First(&reel, reel_id)
	result := db.Delete(&reel).Error
	return result
}
