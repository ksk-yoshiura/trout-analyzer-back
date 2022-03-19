package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Rod struct {
	gorm.Model
	RodHardnessCondition ToolCondition `gorm:"foreignKey:Hardness"`
	Name                 string        `json:"name"`
	UserId               int           `json:"userId"`
	CompanyName          string        `json:"companyName"`
	Length               string        `json:"length"`
	Hardness             string        `json:"hardness"`
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
func GetAllRods(rods []Rod, uid int) []Rod {
	db := database.GetDBConn()
	// ログインユーザは自分のロッドしか見れない
	db.Where("user_id = ?", uid).Preload("RodHardnessCondition").Find(&rods)
	return rods
}

/**
  ロッド取得
*/
func GetRod(rod Rod, rod_id int, uid int) Rod {
	db := database.GetDBConn()
	// ログインユーザは自分のロッドしか見れない
	db.Where("user_id = ?", uid).Preload("RodHardnessCondition").First(&rod, rod_id)
	return rod
}

/**
  ロッド更新
*/
func UpdateRod(r Rod, rod_id int) error {
	var rod Rod
	db := database.GetDBConn()
	// ログインユーザは自分のロッドしか見れない
	db.Where("user_id = ?", r.UserId).First(&rod, rod_id)

	result := db.Model(&rod).Updates(Rod{
		Name:        r.Name,
		UserId:      r.UserId,
		Hardness:    r.Hardness,
		Length:      r.Length,
		CompanyName: r.CompanyName,
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
