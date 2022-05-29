package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/wcl48/valval"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	// HitPattern HitPattern `gorm:"foreignKey:Id;references:RecordId"`
	FieldImage FieldImage `gorm:"foreignKey:FieldId"`
	Field      Field      `gorm:"foreignKey:FieldId"`
	UserId     int        `json:"userId"`
	FieldId    int        `json:"fieldId"`
}

func RecordValidate(record Record) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(record)
}

/**
  レコード一覧取得
*/
func GetAllRecords(records []Record, uid int) []Record {
	db := database.GetDBConn()

	subQueryPatternSum := db.Model(&Record{}).Select("Count(HitPattern.ID)").Preload("HitPattern").Where("user_id = ?", uid)
	subQueryPatternCaughtSum := db.Model(&Record{}).Select("Count(HitPattern.ID)").Preload("HitPattern").Where("user_id = ? AND result = ?", uid, 1)
	subQueryPatternBitSum := db.Model(&Record{}).Select("Count(HitPattern.ID)").Preload("HitPattern").Where("user_id = ? AND result = ?", uid, 2)
	subQueryPatternChasedSum := db.Model(&Record{}).Select("Count(HitPattern.ID)").Preload("HitPattern").Where("user_id = ? AND result = ?", uid, 3)
	// ログインユーザは自分のレコードしか見れない
	db.Table("(?) as pattern_sum, (?) as caught_sum, (?) as bit_sum, (?) as chased_sum", subQueryPatternSum, subQueryPatternCaughtSum, subQueryPatternBitSum, subQueryPatternChasedSum).Where("user_id = ?", uid).Preload("Field.FieldImage").Find(&records)
	return records
}

/**
  レコード取得
*/
func GetRecord(record Record, record_id int, uid int) Record {
	db := database.GetDBConn()
	// ログインユーザは自分のレコードしか見れない
	db.Where("user_id = ?", uid).Preload("Field").First(&record, record_id)
	return record
}

/**
  レコード更新
*/
func UpdateRecord(r Record, record_id int) error {
	var record Record
	db := database.GetDBConn()
	// ログインユーザは自分のレコードしか見れない
	db.Where("user_id = ?", r.UserId).First(&record, record_id)

	result := db.Model(&record).Updates(Record{
		UserId:  r.UserId,
		FieldId: r.FieldId,
	}).Error
	return result
}

/**
  レコード作成
*/
func CreateRecord(record Record) (Record, error) {
	db := database.GetDBConn()
	if result := db.Create(&record); result.Error != nil {
		return record, result.Error
	}
	return record, nil
}

/**
  レコード削除
*/
func DeleteRecord(record Record, record_id int) error {
	db := database.GetDBConn()
	db.First(&record, record_id)
	result := db.Delete(&record).Error
	return result
}
