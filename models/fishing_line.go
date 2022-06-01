package models

import (
	"strconv"
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type FishingLine struct {
	gorm.Model
	LineImage     LineImage     `gorm:"foreignKey:LineId;"`
	LineCondition ToolCondition `gorm:"foreignKey:LineTypeId"`
	Name          string        `json:"name"`
	UserId        int           `json:"userId"`
	LineTypeId    string        `json:"lineTypeId"`
	Thickness     string        `json:"thickness"`
	CompanyName   string        `json:"companyName"`
}

/**
バリデーション
*/
func (line FishingLine) Validate() error {
	return validation.ValidateStruct(&line,
		validation.Field(
			&line.Name,
			validation.Required.Error("Name is required"),
			validation.RuneLength(1, 40).Error("Name should be less thna 40 letters"),
		),
		validation.Field(
			&line.CompanyName,
			validation.RuneLength(0, 80).Error("CompanyName should be less thna 80 letters"),
		),
	)
}

/**
  ライン一覧取得
*/
func GetAllLines(fishing_lines []FishingLine, uid int) []FishingLine {
	db := database.GetDBConn()
	// ログインユーザは自分のラインしか見れない
	db.Where("user_id = ?", uid).Preload("LineCondition").Find(&fishing_lines)
	return fishing_lines
}

/**
  ライン取得
*/
func GetLine(fishing_line FishingLine, line_id int, uid int) FishingLine {
	db := database.GetDBConn()
	// ログインユーザは自分のラインしか見れない
	db.Where("user_id = ?", uid).Preload("LineImage").Preload("LineCondition").First(&fishing_line, line_id)
	return fishing_line
}

/**
  ライン更新
*/
func UpdateLine(f FishingLine, line_id int) error {
	var fishing_line FishingLine
	db := database.GetDBConn()

	db.First(&fishing_line, line_id)

	result := db.Model(&fishing_line).Updates(FishingLine{
		Name:        f.Name,
		UserId:      f.UserId,
		LineTypeId:  f.LineTypeId,
		Thickness:   f.Thickness,
		CompanyName: f.CompanyName,
	}).Error
	return result
}

/**
  ライン作成
*/
func CreateLine(line FishingLine, image Image) error {
	// ライン画像モデル
	var line_image LineImage
	db := database.GetDBConn()
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&line).Error; err != nil {
			// エラーの場合ロールバックされる
			return err
		}
		// 画像データにラインIDをセット
		line_image.LineId = line.ID
		file_name := CreateImageName()
		image_path := "/line_image/" + strconv.Itoa(line.UserId) + "/" + file_name
		line_image.ImageFile = image_path

		if err := tx.Create(&line_image).Error; err != nil {
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
  ライン削除
*/
func DeleteLine(fishing_line FishingLine, line_id int) error {
	db := database.GetDBConn()
	db.First(&fishing_line, line_id)
	result := db.Delete(&fishing_line).Error
	return result
}
