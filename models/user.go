package models

import (
	"github.com/jinzhu/gorm"
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
