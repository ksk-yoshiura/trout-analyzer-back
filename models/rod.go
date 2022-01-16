package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Rod struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	CompanyName string `json:"company_name"`
	Length      string `json:"length"`
	HardnessId  int    `json:"hardness_id"`
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
