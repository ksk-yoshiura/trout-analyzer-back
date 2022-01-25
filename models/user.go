package models

import (
	"regexp"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

/**
  ログイン時ユーザを検索
*/
func FindUser(user User) User {
	db := database.GetDBConn()
	db.Where(user).First(&user)
	return user
}

/**
  ユーザ一覧取得
*/
func GetAllUsers(users []User) []User {
	db := database.GetDBConn()
	db.Find(&users)
	return users
}

/**
  ユーザ取得
*/
func GetUser(user User, uid int) User {
	db := database.GetDBConn()
	db.First(&user, uid)
	return user
}

/**
  ユーザ更新
*/
func UpdateUser(u User, uid int) error {
	var user User
	db := database.GetDBConn()
	db.First(&user, uid)

	result := db.Model(&user).Updates(User{
		Email:     u.Email,
		Password:  u.Password,
		Nickname:  u.Nickname,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		GroupId:   u.GroupId,
	}).Error

	return result
}

/**
  ユーザ作成
*/
func CreateUser(user User) error {
	db := database.GetDBConn()
	result := db.Create(&user).Error
	return result
}

/**
  ユーザ削除
*/
func DeleteUser(user User, uid int) error {
	db := database.GetDBConn()
	db.First(&user, uid)
	result := db.Delete(&user).Error
	return result
}
