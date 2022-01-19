package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
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
func (uc *UsersController) Show(c echo.Context) error {
	user := models.User{}
	uid, _ := strconv.Atoi(c.Param("id"))
	result := models.GetUser(user, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ユーザ更新
*/
func (uc *UsersController) Update(c echo.Context) error {
	user := models.User{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateUser(user, uid, c)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ユーザ作成
*/
func (uc *UsersController) Create(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	result := models.CreateUser(user)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ユーザ削除
*/
func (uc *UsersController) Delete(c echo.Context) error {
	user := models.User{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.DeleteUser(user, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
