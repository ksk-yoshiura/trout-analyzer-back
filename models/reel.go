package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/wcl48/valval"
	"gorm.io/gorm"
)

type Reel struct {
	gorm.Model
	ReelImage           ReelImage     `gorm:"foreignKey:ReelId"`
	GearCondition       ToolCondition `gorm:"foreignKey:Gear"`
	TypeNumberCondition ToolCondition `gorm:"foreignKey:TypeNumber"`
	Name                string        `json:"name"`
	UserId              int           `json:"userId"`
	CompanyName         string        `json:"companyName"`
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
	db.Where("user_id = ?", uid).Preload("ReelImage").Preload("GearCondition").Preload("TypeNumberCondition").First(&reel, reel_id)
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
func CreateReel(reel Reel, image Image) error {
	// リール画像モデル
	var reel_image ReelImage
	db := database.GetDBConn()
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&reel).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}
		// 画像データにリールIDをセット
		reel_image.ReelId = reel.ID
		file_name := CreateImageName()
		image_path := "/reel_image/" + strconv.Itoa(reel.UserId) + "/" + file_name
		reel_image.ImageFile = image_path

		if err := tx.Create(&reel_image).Error; err != nil {
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
  リール削除
*/
func DeleteReel(reel Reel, reel_id int) error {
	db := database.GetDBConn()
	db.First(&reel, reel_id)
	result := db.Delete(&reel).Error
	return result
}
