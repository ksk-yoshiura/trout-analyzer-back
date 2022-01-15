package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DeleteFlg int    `json:"delete_flg"`
	GroupId   int    `json:"group_id"`
}

func UserValidate(user User) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(user)
}
