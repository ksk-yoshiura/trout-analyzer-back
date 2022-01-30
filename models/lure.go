package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Lure struct {
	gorm.Model
	LureType    LureType `gorm:"foreignKey:LureTypeId"`
	Name        string   `json:"name"`
	UserId      int      `json:"user_id"`
	LureTypeId  int      `json:"lure_type_id"`
	CompanyName string   `json:"company_name"`
	Weight      string   `json:"weight"`
	Color       string   `json:"color"`
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
func GetAllLures(lures []Lure, uid int) []Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ?", uid).Preload("LureType").Find(&lures)
	return lures
}

/**
  ルアー取得
*/
func GetLure(lure Lure, lure_id int, uid int) Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ?", uid).Preload("LureType").First(&lure, lure_id)
	return lure
}

/**
  ルアー更新
*/
func UpdateLure(l Lure, lure_id int, uid int) error {
	var lure Lure
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ?", uid).First(&lure, lure_id)

	result := db.Model(&lure).Updates(Lure{
		Name:        l.Name,
		UserId:      l.UserId,
		LureTypeId:  l.LureTypeId,
		CompanyName: l.CompanyName,
		Color:       l.Color,
		Weight:      l.Weight,
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
