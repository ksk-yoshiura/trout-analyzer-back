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

func getUsers(c echo.Context) error {
  users := []User{}
  database.DB.Find(&users)
  return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Take(&user)
  return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Save(&user)
  return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Create(&user)
  return c.JSON(http.StatusCreated, user)
}

func deleteUser(c echo.Context) error {
  id := c.Param("id")
  database.DB.Delete(&User{}, id)
  return c.NoContent(http.StatusNoContent)
}

func newRouter() *echo.Echo {
  e := echo.New()
  database.Connect()
  // sqlDB, _ := database.DB.DB()
  // defer sqlDB.Close()

  e.GET("/users", getUsers)
  e.GET("/users/:id", getUser)
  e.PUT("/users/:id", updateUser)
  e.POST("/users", createUser)
  e.DELETE("/users/:id", deleteUser)
	return e
}