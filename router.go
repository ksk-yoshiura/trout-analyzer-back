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

	// ロッドコントローラ
	rodsController := controllers.NewRodsController()

	e.GET("/rods", rodsController.GetAllRods)
	e.GET("/rods/:id", rodsController.GetRod)
	e.PUT("/rods/:id", rodsController.UpdateRod)
	e.POST("/rods", rodsController.CreateRod)
	e.POST("/rods/:id", rodsController.DeleteRod)

	// リールコントローラ
	reelsController := controllers.NewReelsController()

	e.GET("/reels", reelsController.GetAllReels)
	e.GET("/reels/:id", reelsController.GetReel)
	e.PUT("/reels/:id", reelsController.UpdateReel)
	e.POST("/reels", reelsController.CreateReel)
	e.POST("/reels/:id", reelsController.DeleteReel)

	// タックルコントローラ
	tacklesController := controllers.NewTacklesController()

	e.GET("/tackles", tacklesController.GetAllTackles)
	e.GET("/tackles/:id", tacklesController.GetTackle)
	e.PUT("/tackles/:id", tacklesController.UpdateTackle)
	e.POST("/tackles", tacklesController.CreateTackle)
	e.POST("/tackles/:id", tacklesController.DeleteTackle)

	// フィールドコントローラ
	fieldsController := controllers.NewFieldsController()

	e.GET("/fields", fieldsController.GetAllFields)
	e.GET("/fields/:id", fieldsController.GetField)
	e.PUT("/fields/:id", fieldsController.UpdateField)
	e.POST("/fields", fieldsController.CreateField)
	e.POST("/fields/:id", fieldsController.DeleteField)

	// ラインコントローラ
	fishingLinesController := controllers.NewFishingLinesController()

	e.GET("/lines", fishingLinesController.GetAllFishingLines)
	e.GET("/lines/:id", fishingLinesController.GetFishingLine)
	e.PUT("/lines/:id", fishingLinesController.UpdateFishingLine)
	e.POST("/lines", fishingLinesController.CreateFishingLine)
	e.POST("/lines/:id", fishingLinesController.DeleteFishingLine)
	return e
}
