package models

import (
	"regexp"
	"strconv"
	"trout-analyzer-back/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
func UpdateUser(user User, uid int, c echo.Context) error {
	db := database.GetDBConn()

	db.First(&user, uid)
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	nickname := c.FormValue("nickname")
	firstname := c.FormValue("first_name")
	lastname := c.FormValue("last_name")
	groupid, _ := strconv.Atoi(c.FormValue("group_id"))

	result := db.Model(&user).Updates(User{
		Name:      name,
		Email:     email,
		Password:  password,
		Nickname:  nickname,
		FirstName: firstname,
		LastName:  lastname,
		GroupId:   groupid,
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
