package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
)

// UsersController controller for Users request
type UsersController struct{}

// NewUsersController is constructer for UsersController
func NewUsersController() *UsersController {
	return new(UsersController)
}

/**
  ユーザ一覧取得
*/
func (uc *UsersController) Index(c echo.Context) error {
	users := []models.User{}
	result := models.GetAllUsers(users)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ユーザ取得
*/
func (uc *UsersController) GetUser(c echo.Context) error {
	db := database.GetDBConn()
	user := models.User{}
	uid, _ := strconv.Atoi(c.Param("id"))
	db.First(&user, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		user,
	))
}

/**
  ユーザ更新
*/
func (uc *UsersController) UpdateUser(c echo.Context) error {
	db := database.GetDBConn()
	user := models.User{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&user, uid)
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	nickname := c.FormValue("nickname")
	firstname := c.FormValue("first_name")
	lastname := c.FormValue("last_name")
	groupid, _ := strconv.Atoi(c.FormValue("group_id"))

	db.Model(&user).Updates(models.User{
		Name:      name,
		Email:     email,
		Password:  password,
		Nickname:  nickname,
		FirstName: firstname,
		LastName:  lastname,
		GroupId:   groupid,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		user,
	))
}

/**
  ユーザ作成
*/
func (uc *UsersController) CreateUser(c echo.Context) error {
	db := database.GetDBConn()
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	result := db.Create(&user).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ユーザ削除
*/
func (uc *UsersController) DeleteUser(c echo.Context) error {
	db := database.GetDBConn()
	user := models.User{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&user, uid)
	result := db.Delete(&user)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
