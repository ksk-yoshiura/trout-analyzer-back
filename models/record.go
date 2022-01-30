package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Record struct {
	gorm.Model
	Field   Field `gorm:"foreignKey:FieldId"`
	UserId  int   `json:"user_id"`
	FieldId int   `json:"field_id"`
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
	// ログインユーザは自分のレコードしか見れない
	db.Where("user_id = ?", uid).Preload("Field").Find(&records)
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
func CreateRecord(record Record) error {
	db := database.GetDBConn()
	result := db.Create(&record).Error
	return result
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
