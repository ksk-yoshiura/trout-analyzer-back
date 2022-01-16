package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Lure struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	LureTypeId  int    `json:"lure_type_id"`
	CompanyName string `json:"company_name"`
	Weight      string `json:"weight"`
	Color       string `json:"color"`
}

func LureValidate(lure Lure) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(lure)
}
