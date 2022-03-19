package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Reel struct {
	gorm.Model
	GearCondition       ToolCondition `gorm:"foreignKey:Gear"`
	TypeNumberCondition ToolCondition `gorm:"foreignKey:TypeNumber"`
	Name                string        `json:"name"`
	UserId              int           `json:"userId"`
	CompanyName         string        `json:"company"`
	TypeNumber          string        `json:"typeNumberId"`
	Gear                string        `json:"gearId"`
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
func GetAllReels(reels []Reel, uid int) []Reel {
	db := database.GetDBConn()
	// ログインユーザは自分のリールしか見れない
	db.Where("user_id = ?", uid).Preload("GearCondition").Preload("TypeNumberCondition").Find(&reels)
	return reels
}

/**
  リール取得
*/
func GetReel(reel Reel, reel_id int, uid int) Reel {
	db := database.GetDBConn()
	// ログインユーザは自分のリールしか見れない
	db.Where("user_id = ?", uid).Preload("GearCondition").Preload("TypeNumberCondition").First(&reel, reel_id)
	return reel
}

/**
  リール更新
*/
func UpdateReel(r Reel, reel_id int) error {
	var reel Reel
	db := database.GetDBConn()
	db.First(&reel, reel_id)

	result := db.Model(&reel).Updates(Reel{
		Name:        r.Name,
		UserId:      r.UserId,
		TypeNumber:  r.TypeNumber,
		Gear:        r.Gear,
		CompanyName: r.CompanyName,
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
