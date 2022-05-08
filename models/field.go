package models

import (
	"regexp"
	"time"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Field struct {
	gorm.Model
	Name          string    `json:"name"`
	UserId        int       `json:"user_id"`
	Address       string    `json:"address"`
	LastVisitedAt time.Time `json:"last_visited_at"`
}

func FieldValidate(field Field) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(field)
}

/**
  フィールド一覧取得
*/
func GetAllFields(fields []Field, uid int) []Field {
	db := database.GetDBConn()
	// ログインユーザは自分のフィールドしか見れない
	db.Where("user_id = ?", uid).Find(&fields)
	return fields
}

/**
  フィールド取得
*/
func GetField(field Field, field_id int, uid int) Field {
	db := database.GetDBConn()
	// ログインユーザは自分のフィールドしか見れない
	db.Where("user_id = ?", uid).First(&field, field_id)
	return field
}

/**
  フィールド更新
*/
func UpdateField(f Field, field_id int) error {
	var field Field
	db := database.GetDBConn()
	// ログインユーザは自分のフィールドしか見れない
	db.Where("user_id = ?", f.UserId).First(&field, field_id)

	result := db.Model(&field).Updates(Field{
		Name:    f.Name,
		UserId:  f.UserId,
		Address: f.Address,
	}).Error
	return result
}

/**
  フィールド作成
*/
func CreateField(field Field) error {
	db := database.GetDBConn()
	result := db.Create(&field).Error
	return result
}

/**
  フィールド削除
*/
func DeleteField(field Field, field_id int) error {
	db := database.GetDBConn()
	db.First(&field, field_id)
	result := db.Delete(&field).Error
	return result
}

/**
  フィールド訪問最終日更新
*/
func RecordLastVisitDate(uid int, field_id int) error {
	var field Field
	db := database.GetDBConn()
	// ログインユーザは自分のフィールドしか見れない
	db.Where("user_id = ?", uid).First(&field, field_id)

	result := db.Model(&field).Updates(Field{
		LastVisitedAt: time.Now(),
	}).Error
	return result
}
