package models

import (
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Field   Field `gorm:"foreignKey:FieldId"`
	UserId  int   `json:"userId"`
	FieldId int   `json:"fieldId"`
}

/**
バリデーション
*/
func (record Record) Validate() error {
	return validation.ValidateStruct(&record,
		validation.Field(
			&record.FieldId,
			validation.Required.Error("Field is required"),
		),
	)
}

/**
  レコード一覧取得
*/
func GetAllRecords(records []Record, uid int) []Record {
	db := database.GetDBConn()

	// ログインユーザは自分のレコードしか見れない
	db.Preload("Field.FieldImage").Find(&records)
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
