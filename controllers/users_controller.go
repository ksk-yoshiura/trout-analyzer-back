package controllers

import (
	"net/http"

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
	u := models.User{}
	result := u.GetUsers()

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

// /**
//   ユーザ取得
// */
// func getUser(c echo.Context) error {
// 	user := User{}
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	database.DB.Take(&user)
// 	return c.JSON(http.StatusOK, user)
// }

// /**
//   ユーザ更新
// */
// func updateUser(c echo.Context) error {
// 	user := User{}
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	database.DB.Save(&user)
// 	return c.JSON(http.StatusOK, user)
// }

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
