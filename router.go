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

	// ユーザコントローラ
	usersController := controllers.NewUsersController()

	e.GET("/users", usersController.GetAllUsers)
	e.GET("/users/:id", usersController.GetUser)
	e.PUT("/users/:id", usersController.UpdateUser)
	e.POST("/users", usersController.CreateUser)
	e.POST("/users/:id", usersController.DeleteUser)

	// ルアーコントローラ
	luresController := controllers.NewLuresController()

	e.GET("/lures", luresController.GetAllLures)
	e.GET("/lures/:id", luresController.GetLure)
	e.PUT("/lures/:id", luresController.UpdateLure)
	e.POST("/lures", luresController.CreateLure)
	e.POST("/lures/:id", luresController.DeleteLure)
	return e
}
