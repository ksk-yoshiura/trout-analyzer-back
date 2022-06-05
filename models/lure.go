package models

import (
	"strconv"
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Lure struct {
	gorm.Model
	LureImage   LureImage `gorm:"foreignKey:LureId"`
	LureType    LureType  `gorm:"foreignKey:LureTypeId"`
	Color       Color     `gorm:"foreignKey:ColorId"`
	Name        string    `json:"name"`
	UserId      int       `json:"userId"`
	LureTypeId  string    `json:"lureTypeId"`
	CompanyName string    `json:"companyName"`
	Weight      string    `json:"weight"`
	ColorId     int       `json:"color"`
}

/**
バリデーション
*/
func (lure Lure) Validate() error {
	return validation.ValidateStruct(&lure,
		validation.Field(
			&lure.Name,
			validation.Required.Error("Name is required"),
			validation.RuneLength(1, 40).Error("Name should be less thna 40 letters"),
		),
		validation.Field(
			&lure.ColorId,
			validation.Required.Error("Color is required"),
		),
		validation.Field(
			&lure.LureTypeId,
			validation.Required.Error("LureType is required"),
		),
		validation.Field(
			&lure.CompanyName,
			validation.RuneLength(1, 80).Error("CompanyName should be less thna 80 letters"),
		),
	)
}

/**
  ルアー一覧取得
*/
func GetAllLures(lures []Lure, uid int) []Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ?", uid).Preload("LureImage").Preload("LureType").Find(&lures)
	return lures
}

/**
  ルアータイプ毎のルアー一覧取得
*/
func GetLuresSelectedLureType(lures []Lure, type_id string, uid int) []Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ? AND lure_type_id = ?", uid, type_id).Preload("LureImage").Preload("Color").Preload("LureType").Find(&lures)
	return lures
}

/**
  ルアー取得
*/
func GetLure(lure Lure, lure_id int, uid int) Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ?", uid).Preload("Color").Preload("LureImage").Preload("LureType").First(&lure, lure_id)
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
		ColorId:     l.ColorId,
		Weight:      l.Weight,
	}).Error
	return result
}

/**
  ルアー作成
*/
func CreateLure(lure Lure, image Image) error {
	// ルアー画像モデル
	var lure_image LureImage
	db := database.GetDBConn()
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&lure).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}
		// 画像データにルアーIDをセット
		lure_image.LureId = lure.ID
		file_name := CreateImageName()
		image_path := "/lure_image/" + strconv.Itoa(lure.UserId) + "/" + file_name
		lure_image.ImageFile = image_path

		if err := tx.Create(&lure_image).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}

		// S3に画像アップロード
		UploadToS3(image, image_path)
		// nilが返却されるとトランザクション内の全処理がコミットされる
		return nil
	})
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
