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
	database.Connect()
	// sqlDB, _ := database.DB.DB()
	// defer sqlDB.Close()

	usersController := controllers.NewUsersController()

	e.GET("/users", usersController.Index)
	// e.GET("/user/:id", getUser)
	// e.PUT("/user/:id", updateUser)
	// e.POST("/user", createUser)
	// e.DELETE("/user/:id", deleteUser)
	return e
}
