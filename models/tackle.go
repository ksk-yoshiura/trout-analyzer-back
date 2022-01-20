package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/wcl48/valval"
)

type Tackle struct {
	gorm.Model
	UserId int `json:"user_id"`
	RodId  int `json:"rod_id"`
	ReelId int `json:"reel_id"`
	LineId int `json:"liner_id"`
}

func TackleValidate(tackle Tackle) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(tackle)
}

/**
  タックル一覧取得
*/
func GetAllTackles(tackles []Tackle) []Tackle {
	db := database.GetDBConn()
	db.Find(&tackles)
	return tackles
}

/**
  タックル取得
*/
func GetTackle(tackle Tackle, tackle_id int) Tackle {
	db := database.GetDBConn()
	db.First(&tackle, tackle_id)
	return tackle
}

/**
  タックル更新
*/
func UpdateTackle(tackle Tackle, tackle_id int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&tackle, tackle_id)
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	rod_id, _ := strconv.Atoi(c.FormValue("rod_id"))
	reel_id, _ := strconv.Atoi(c.FormValue("reel_id"))
	line_id, _ := strconv.Atoi(c.FormValue("line_id"))

	result := db.Model(&tackle).Updates(Tackle{
		UserId: user_id,
		RodId:  rod_id,
		ReelId: reel_id,
		LineId: line_id,
	}).Error
	return result
}

/**
  タックル作成
*/
func CreateTackle(tackle Tackle) error {
	db := database.GetDBConn()
	result := db.Create(&tackle).Error
	return result
}

/**
  タックル削除
*/
func DeleteTackle(tackle Tackle, tackle_id int) error {
	db := database.GetDBConn()
	db.First(&tackle, tackle_id)
	result := db.Delete(&tackle).Error
	return result
}
