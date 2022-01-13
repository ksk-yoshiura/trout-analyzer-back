package main

import (
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	// "github.com/x-color/simple-webapp/handler"
	"trout-analyzer-back/controllers"
	"trout-analyzer-back/database"
)

/**
  ルート定義
*/
func newRouter() *echo.Echo {
	e := echo.New()
	database.GetDBConfig()

	usersController := controllers.NewUsersController()

	e.GET("/users", usersController.GetAllUsers)
	e.GET("/users/:id", usersController.GetUser)
	e.POST("/users/:id", usersController.UpdateUser)
	e.POST("/users", usersController.CreateUser)
	return e
}
