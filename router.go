package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// "github.com/x-color/simple-webapp/handler"
	"trout-analyzer-back/controllers"
)

/**
  ルート定義
*/
func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// サインアップ
	e.POST("/signup", controllers.Signup)
	// ログイン
	e.POST("/login", controllers.Login)

	// /api 下はJWTの認証が必要
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(controllers.Config))
	// ユーザコントローラー
	usersController := controllers.NewUsersController()

	api.GET("/users", usersController.Index)
	api.GET("/user", usersController.Show)
	api.PUT("/user", usersController.Update)
	api.POST("/user", usersController.Create)
	api.POST("/users/:id", usersController.Delete)

	// ルアーコントローラー
	luresController := controllers.NewLuresController()

	api.GET("/lures", luresController.Index)
	api.GET("/lures/:id", luresController.Show)
	api.PUT("/lures/:id", luresController.Update)
	api.POST("/lures", luresController.Create)
	api.POST("/lures/:id", luresController.Delete)

	// ロッドコントローラー
	rodsController := controllers.NewRodsController()

	api.GET("/rods", rodsController.Index)
	api.GET("/rods/:id", rodsController.Show)
	api.PUT("/rods/:id", rodsController.Update)
	api.POST("/rods", rodsController.Create)
	api.POST("/rods/:id", rodsController.Delete)

	// リールコントローラー
	reelsController := controllers.NewReelsController()

	api.GET("/reels", reelsController.Index)
	api.GET("/reels/:id", reelsController.Show)
	api.PUT("/reels/:id", reelsController.Update)
	api.POST("/reels", reelsController.Create)
	api.POST("/reels/:id", reelsController.Delete)

	// タックルコントローラー
	tacklesController := controllers.NewTacklesController()

	api.GET("/tackles", tacklesController.Index)
	api.GET("/tackles/:id", tacklesController.Show)
	api.PUT("/tackles/:id", tacklesController.Update)
	api.POST("/tackles", tacklesController.Create)
	api.POST("/tackles/:id", tacklesController.Delete)

	// フィールドコントローラー
	fieldsController := controllers.NewFieldsController()

	api.GET("/fields", fieldsController.Index)
	api.GET("/fields/:id", fieldsController.Show)
	api.PUT("/fields/:id", fieldsController.Update)
	api.POST("/fields", fieldsController.Create)
	api.POST("/fields/:id", fieldsController.Delete)

	// ラインコントローラー
	fishingLinesController := controllers.NewFishingLinesController()

	api.GET("/lines", fishingLinesController.Index)
	api.GET("/lines/:id", fishingLinesController.Show)
	api.PUT("/lines/:id", fishingLinesController.Update)
	api.POST("/lines", fishingLinesController.Create)
	api.POST("/lines/:id", fishingLinesController.Delete)

	// ヒットパターンコントローラー
	hitPatternsController := controllers.NewHitPatternsController()

	api.GET("/patterns", hitPatternsController.Index)
	api.GET("/patterns/:id", hitPatternsController.Show)
	api.PUT("/patterns/:id", hitPatternsController.Update)
	api.POST("/patterns", hitPatternsController.Create)
	api.POST("/patterns/:id", hitPatternsController.Delete)
	return e
}
