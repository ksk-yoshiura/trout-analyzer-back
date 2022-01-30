package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Tackle struct {
	gorm.Model
	Reel        Reel        `gorm:"foreignKey:ReelId"`
	Rod         Rod         `gorm:"foreignKey:RodId"`
	FishingLine FishingLine `gorm:"foreignKey:LineId"`
	UserId      int         `json:"user_id"`
	RodId       int         `json:"rod_id"`
	ReelId      int         `json:"reel_id"`
	LineId      int         `json:"liner_id"`
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
func GetAllTackles(tackles []Tackle, uid int) []Tackle {
	db := database.GetDBConn()
	// ログインユーザは自分のタックルしか見れない
	db.Where("user_id = ?", uid).Preload("Rod").Preload("Reel").Preload("FishingLine").Find(&tackles)
	return tackles
}

/**
  タックル取得
*/
func GetTackle(tackle Tackle, tackle_id int, uid int) Tackle {
	db := database.GetDBConn()
	// ログインユーザは自分のタックルしか見れない
	db.Where("user_id = ?", uid).Preload("Rod").Preload("Reel").Preload("FishingLine").First(&tackle, tackle_id)
	return tackle
}

/**
  タックル更新
*/
func UpdateTackle(t Tackle, tackle_id int) error {
	var tackle Tackle
	db := database.GetDBConn()

	// ログインユーザは自分のタックルしか見れない
	db.Where("user_id = ?", t.UserId).First(&tackle, tackle_id)

	result := db.Model(&tackle).Updates(Tackle{
		UserId: t.UserId,
		RodId:  t.RodId,
		ReelId: t.ReelId,
		LineId: t.LineId,
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
