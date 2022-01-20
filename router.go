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

	// ユーザコントローラー
	usersController := controllers.NewUsersController()

	e.GET("/users", usersController.Index)
	e.GET("/users/:id", usersController.Show)
	e.PUT("/users/:id", usersController.Update)
	e.POST("/users", usersController.Create)
	e.POST("/users/:id", usersController.Delete)

	// ルアーコントローラー
	luresController := controllers.NewLuresController()

	e.GET("/lures", luresController.Index)
	e.GET("/lures/:id", luresController.Show)
	e.PUT("/lures/:id", luresController.Update)
	e.POST("/lures", luresController.CreateLure)
	e.POST("/lures/:id", luresController.DeleteLure)

	// ロッドコントローラー
	rodsController := controllers.NewRodsController()

	e.GET("/rods", rodsController.GetAllRods)
	e.GET("/rods/:id", rodsController.GetRod)
	e.PUT("/rods/:id", rodsController.UpdateRod)
	e.POST("/rods", rodsController.CreateRod)
	e.POST("/rods/:id", rodsController.DeleteRod)

	// リールコントローラー
	reelsController := controllers.NewReelsController()

	e.GET("/reels", reelsController.Index)
	e.GET("/reels/:id", reelsController.Show)
	e.PUT("/reels/:id", reelsController.Update)
	e.POST("/reels", reelsController.Create)
	e.POST("/reels/:id", reelsController.Delete)

	// タックルコントローラー
	tacklesController := controllers.NewTacklesController()

	e.GET("/tackles", tacklesController.GetAllTackles)
	e.GET("/tackles/:id", tacklesController.GetTackle)
	e.PUT("/tackles/:id", tacklesController.UpdateTackle)
	e.POST("/tackles", tacklesController.CreateTackle)
	e.POST("/tackles/:id", tacklesController.DeleteTackle)

	// フィールドコントローラー
	fieldsController := controllers.NewFieldsController()

	e.GET("/fields", fieldsController.Index)
	e.GET("/fields/:id", fieldsController.Show)
	e.PUT("/fields/:id", fieldsController.Update)
	e.POST("/fields", fieldsController.Create)
	e.POST("/fields/:id", fieldsController.Delete)

	// ラインコントローラー
	fishingLinesController := controllers.NewFishingLinesController()

	e.GET("/lines", fishingLinesController.GetAllFishingLines)
	e.GET("/lines/:id", fishingLinesController.GetFishingLine)
	e.PUT("/lines/:id", fishingLinesController.UpdateFishingLine)
	e.POST("/lines", fishingLinesController.CreateFishingLine)
	e.POST("/lines/:id", fishingLinesController.DeleteFishingLine)

	// ヒットパターンコントローラー
	hitPatternsController := controllers.NewHitPatternsController()

	e.GET("/patterns", hitPatternsController.GetAllHitPatterns)
	e.GET("/patterns/:id", hitPatternsController.GetHitPattern)
	e.PUT("/patterns/:id", hitPatternsController.UpdateHitPattern)
	e.POST("/patterns", hitPatternsController.CreateHitPattern)
	e.POST("/patterns/:id", hitPatternsController.DeleteHitPattern)
	return e
}
