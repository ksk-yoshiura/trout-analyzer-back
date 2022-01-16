package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type FishingLine struct {
	gorm.Model
	Name        string `json:"name"`
	UserId      int    `json:"user_id"`
	LineTypeId  int    `json:"line_type_id"`
	Thickness   int    `json:"thickness"`
	CompanyName string `json:"company_name"`
}

func FishingLineValidate(fishing_line FishingLine) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(fishing_line)
}
