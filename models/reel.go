package models

import (
	"strconv"
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
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

/**
バリデーション
*/
func (reel Reel) Validate() error {
	return validation.ValidateStruct(&reel,
		validation.Field(
			&reel.Name,
			validation.Required.Error("Name is required"),
			validation.RuneLength(1, 40).Error("Name should be less thna 40 letters"),
		),
		validation.Field(
			&reel.CompanyName,
			validation.RuneLength(1, 80).Error("CompanyName should be less thna 80 letters"),
		),
	)
}

/**
  リール一覧取得
*/
func GetAllReels(reels []Reel, uid int) []Reel {
	db := database.GetDBConn()
	// ログインユーザは自分のリールしか見れない
	db.Where("user_id = ?", uid).Preload("ReelImage").Preload("GearCondition").Preload("TypeNumberCondition").Find(&reels)
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
func UpdateReel(reel Reel, reel_id int, image Image) error {
	// リール画像モデル
	var reel_image ReelImage
	db := database.GetDBConn()

	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND id = ?", reel.UserId, reel_id).Updates(&reel).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}

		if (Image{}) != image { // 画像データがセットされている場合
			// 画像データにリールIDをセット
			reel_image.ReelId = reel.ID
			file_name := CreateImageName()
			image_path := "/reel_image/" + strconv.Itoa(reel.UserId) + "/" + file_name
			reel_image.ImageFile = image_path

			if err := tx.Where("reel_id = ?", reel_id).Updates(&reel_image).Error; err != nil {
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

		if (Image{}) != image { // 画像データがセットされている場合
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
		}
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
