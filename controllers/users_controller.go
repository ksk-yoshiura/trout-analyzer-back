package controllers

import (
	"net/http"

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
func (uc *UsersController) GetAllUsers(c echo.Context) error {
	db := database.GetDBConn()
	u := models.User{}
	db.Find(&u)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		u,
	))
}

/**
  ユーザ取得
*/
func (uc *UsersController) GetUser(c echo.Context) error {
	db := database.GetDBConn()
	user := models.User{}
	uid := c.Param("id")
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
	uid := c.Param("id")
	db.First(&user, uid)
	user.Name = c.FormValue("name")
	db.Update(&user)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		user,
	))
}

// /**
//   ユーザ作成
// */
// func createUser(c echo.Context) error {
// 	user := User{}
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	database.DB.Create(&user)
// 	return c.JSON(http.StatusCreated, newResponse(
// 		http.StatusOK,
// 		http.StatusText(http.StatusOK),
// 		"OK",
// 	))
// }

// /**
//   ユーザ削除
// */
// func deleteUser(c echo.Context) error {
// 	id := c.Param("id")
// 	database.DB.Delete(&User{}, id)
// 	return c.NoContent(http.StatusNoContent)
// }
