package models

import (
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Tackle struct {
	gorm.Model
	Reel   Reel        `gorm:"foreignKey:ReelId"`
	Rod    Rod         `gorm:"foreignKey:RodId"`
	Line   FishingLine `gorm:"foreignKey:LineId"`
	UserId int         `json:"userId"`
	RodId  int         `json:"rodId"`
	ReelId int         `json:"reelId"`
	LineId int         `json:"lineId"`
}

/**
バリデーション
*/
func (tackle Tackle) Validate() error {
	return validation.ValidateStruct(&tackle,
		validation.Field(
			&tackle.RodId,
			validation.Required.Error("Rod is required"),
		),
		validation.Field(
			&tackle.ReelId,
			validation.Required.Error("Reel is required"),
		),
		validation.Field(
			&tackle.LineId,
			validation.Required.Error("Line is required"),
		),
	)
}

/**
  タックル一覧取得
*/
func GetAllTackles(tackles []Tackle, uid int) []Tackle {
	db := database.GetDBConn()
	// ログインユーザは自分のタックルしか見れない
	db.Where("user_id = ?", uid).Preload("Rod.RodHardnessCondition").Preload("Reel.TypeNumberCondition").Preload("Line.LineCondition").Find(&tackles)
	return tackles
}

/**
  タックル取得
*/
func GetTackle(tackle Tackle, tackle_id int, uid int) Tackle {
	db := database.GetDBConn()
	// ログインユーザは自分のタックルしか見れない
	db.Where("user_id = ?", uid).Preload("Rod.RodImage").Preload("Reel.ReelImage").Preload("Line.LineImage").First(&tackle, tackle_id)
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
