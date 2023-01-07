package models

import (
	"strconv"
	"time"
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Field struct {
	gorm.Model
	FieldImage    FieldImage `gorm:"foreignKey:FieldId"`
	Name          string     `json:"name"`
	UserId        int        `json:"user_id"`
	Address       string     `json:"address"`
	LastVisitedAt time.Time  `json:"lastVisitedAt"`
}

/**
バリデーション
*/
func (field Field) Validate() error {
	return validation.ValidateStruct(&field,
		validation.Field(
			&field.Name,
			validation.Required.Error("Name is required"),
			validation.RuneLength(1, 40).Error("Name should be less thna 40 letters"),
		),
		validation.Field(
			&field.Address,
			validation.Required.Error("Address is required"),
			validation.RuneLength(1, 80).Error("Address should be less thna 80 letters"),
		),
	)
}

/**
  フィールド一覧取得
*/
func GetAllFields(fields []Field, uid int) []Field {
	db := database.GetDBConn()
	// ログインユーザは自分のフィールドしか見れない
	db.Where("user_id = ?", uid).Preload("FieldImage").Find(&fields)
	return fields
}

/**
  フィールド取得
*/
func GetField(field Field, field_id int, uid int) Field {
	db := database.GetDBConn()
	// ログインユーザは自分のフィールドしか見れない
	db.Where("user_id = ?", uid).Preload("FieldImage").First(&field, field_id)
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
func CreateField(field Field, image Image) error {
	// フィールド画像モデル
	var field_image FieldImage
	db := database.GetDBConn()
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&field).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}

		if (Image{}) != image { // 画像データがセットされている場合
			// 画像データにフィールドIDをセット
			field_image.FieldId = field.ID
			file_name := CreateImageName()
			image_path := "/field_image/" + strconv.Itoa(field.UserId) + "/" + file_name
			field_image.ImageFile = image_path

			if err := tx.Create(&field_image).Error; err != nil {
				// エラーの場合ロールバックされる
				return err
			}

			// S3に画像アップロード
			UploadToS3(image, image_path)
		}
		// nilが返却されるとトランザクション内の全処理がコミットされる
		return nil
	})
	return result
}

/**
  フィールド削除
*/
func DeleteField(field Field, field_id int) error {
	db := database.GetDBConn()
	db.First(&field, field_id)
	result := db.Select("FieldImage").Delete(&field).Error
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
