package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Reel struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	CompanyName string `json:"company_name"`
	TypeNumber  string `json:"type_number"`
	Gear        string `json:"gear"`
	DeleteFlg   int    `json:"delete_flg"`
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
