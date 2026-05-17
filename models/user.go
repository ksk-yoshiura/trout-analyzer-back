package models

import (
	"trout-analyzer-back/database"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"mailaddress"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	GroupId   int    `json:"groupId"`
}

type NewPassword struct {
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

/**
バリデーション
*/
func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required"),
		),
		validation.Field(
			&user.Password,
			validation.RuneLength(8, 80).Error("Password should be more than 8 letters"),
			validation.Required.Error("Password is required"),
		),
	)
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
