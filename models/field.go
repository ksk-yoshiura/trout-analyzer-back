package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Field struct {
	gorm.Model
	Name    string `json:"name"`
	UserId  int    `json:"user_id"`
	Address string `json:"address"`
}

func FieldValidate(field Field) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(field)
}
