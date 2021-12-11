package main

import (
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	// "github.com/x-color/simple-webapp/handler"
	"trout-analyzer-back/database"
  "net/http"
)

type User struct {
  Id    int    `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

/**
  ユーザ一覧取得
*/
func getUsers(c echo.Context) error {
  users := []User{}
  database.DB.Find(&users)
  return c.JSON(http.StatusOK, users)
}

/** 
  ユーザ取得
*/
func getUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Take(&user)
  return c.JSON(http.StatusOK, user)
}

/** 
  ユーザ更新
*/
func updateUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Save(&user)
  return c.JSON(http.StatusOK, user)
}

/** 
  ユーザ作成
*/
func createUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Create(&user)
  return c.JSON(http.StatusCreated, user)
}

/** 
  ユーザ削除
*/
func deleteUser(c echo.Context) error {
  id := c.Param("id")
  database.DB.Delete(&User{}, id)
  return c.NoContent(http.StatusNoContent)
}

/** 
  ルート定義
*/
func newRouter() *echo.Echo {
  e := echo.New()
  database.Connect()
  // sqlDB, _ := database.DB.DB()
  // defer sqlDB.Close()

  e.GET("/users", getUsers)
  e.GET("/user/:id", getUser)
  e.PUT("/user/:id", updateUser)
  e.POST("/user", createUser)
  e.DELETE("/user/:id", deleteUser)
	return e
}