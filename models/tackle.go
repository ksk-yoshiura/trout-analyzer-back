package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Tackle struct {
	gorm.Model
	Name      string `json:"name"`
	UserId    int    `json:"user_id"`
	RodId     int    `json:"rod_id"`
	ReelId    int    `json:"reel_id"`
	LineId    int    `json:"liner_id"`
	DeleteFlg int    `json:"delete_flg"`
}

func TackleValidate(tackle Tackle) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(tackle)
}