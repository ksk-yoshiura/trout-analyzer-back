package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type HitPattern struct {
	gorm.Model
	UserId   int `json:"user_id"`
	LureId   int `json:"lure_id"`
	TackleId int `json:"tackle_id"`
	Speed    int `json:"speed"`
	Depth    int `json:"depth"`
	Weather  int `json:"weather"`
	Result   int `json:"result"`
	FieldId  int `json:"field_id"`
}

func HitPatternValidate(hit_pattern HitPattern) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(hit_pattern)
}
