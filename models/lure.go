package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/wcl48/valval"
	"gorm.io/gorm"
)

type Lure struct {
	gorm.Model
	LureImage   LureImage `gorm:"foreignKey:LureId"`
	LureType    LureType  `gorm:"foreignKey:LureTypeId"`
	Name        string    `json:"name"`
	UserId      int       `json:"userId"`
	LureTypeId  string    `json:"lureTypeId"`
	CompanyName string    `json:"companyName"`
	Weight      string    `json:"weight"`
	Color       string    `json:"color"`
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

func GetLuresSelectedLureType(lures []Lure, type_id string, uid int) []Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ? AND lure_type_id = ?", uid, type_id).Preload("LureType").Find(&lures)
	return lures
}

/**
  ルアー取得
*/
func GetLure(lure Lure, lure_id int, uid int) Lure {
	db := database.GetDBConn()
	// ログインユーザは自分のルアーしか見れない
	db.Where("user_id = ?", uid).Preload("LureImage").Preload("LureType").First(&lure, lure_id)
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
