package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Rod struct {
	gorm.Model
	RodImage             RodImage      `gorm:"foreignKey:RodId"`
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
	db.Where("user_id = ?", uid).Preload("RodImage").Preload("RodHardnessCondition").First(&rod, rod_id)
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
func CreateRod(rod Rod, image Image) error {
	// ロッド画像モデル
	var rod_image RodImage
	db := database.GetDBConn()
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rod).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}
		// 画像データにロッドIDをセット
		rod_image.RodId = rod.ID
		file_name := CreateImageName()
		image_path := "/rod_image/" + strconv.Itoa(rod.UserId) + "/" + file_name
		rod_image.ImageFile = image_path

		if err := tx.Create(&rod_image).Error; err != nil {
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
  ロッド削除
*/
func DeleteRod(rod Rod, rod_id int) error {
	db := database.GetDBConn()
	db.First(&rod, rod_id)
	result := db.Delete(&rod).Error
	return result
}
