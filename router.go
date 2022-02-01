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
	api.GET("/lure/:id", luresController.Show)
	api.PUT("/lure/:id", luresController.Update)
	api.POST("/lure", luresController.Create)
	api.POST("/lure/:id", luresController.Delete)

	// ロッドコントローラー
	rodsController := controllers.NewRodsController()

	api.GET("/rods", rodsController.Index)
	api.GET("/rod/:id", rodsController.Show)
	api.PUT("/rod/:id", rodsController.Update)
	api.POST("/rod", rodsController.Create)
	api.POST("/rod/:id", rodsController.Delete)

	// リールコントローラー
	reelsController := controllers.NewReelsController()

	api.GET("/reels", reelsController.Index)
	api.GET("/reel/:id", reelsController.Show)
	api.PUT("/reel/:id", reelsController.Update)
	api.POST("/reel", reelsController.Create)
	api.POST("/reel/:id", reelsController.Delete)

	// タックルコントローラー
	tacklesController := controllers.NewTacklesController()

	api.GET("/tackles", tacklesController.Index)
	api.GET("/tackle/:id", tacklesController.Show)
	api.PUT("/tackle/:id", tacklesController.Update)
	api.POST("/tackle", tacklesController.Create)
	api.POST("/tackle/:id", tacklesController.Delete)

	// フィールドコントローラー
	fieldsController := controllers.NewFieldsController()

	api.GET("/fields", fieldsController.Index)
	api.GET("/field/:id", fieldsController.Show)
	api.PUT("/field/:id", fieldsController.Update)
	api.POST("/field", fieldsController.Create)
	api.POST("/field/:id", fieldsController.Delete)

	// ラインコントローラー
	fishingLinesController := controllers.NewFishingLinesController()

	api.GET("/lines", fishingLinesController.Index)
	api.GET("/line/:id", fishingLinesController.Show)
	api.PUT("/line/:id", fishingLinesController.Update)
	api.POST("/line", fishingLinesController.Create)
	api.POST("/line/:id", fishingLinesController.Delete)

	// レコードコントローラー
	recordsController := controllers.NewRecordsController()

	api.GET("/records", recordsController.Index)
	api.GET("/record/:id", recordsController.Show)
	api.PUT("/record/:id", recordsController.Update)
	api.POST("/record", recordsController.Create)
	api.POST("/record/:id", recordsController.Delete)

	// ヒットパターンコントローラー
	hitPatternsController := controllers.NewHitPatternsController()

	api.GET("/patterns", hitPatternsController.Index)
	api.GET("/pattern/:id", hitPatternsController.Show)
	api.PUT("/pattern/:id", hitPatternsController.Update)
	api.POST("/pattern", hitPatternsController.Create)
	api.POST("/pattern/:id", hitPatternsController.Delete)

	// ルアータイプコントローラー
	LureTypesController := controllers.NewLureTypesController()

	api.GET("/lure_types", LureTypesController.Index)

	// ツール条件各種コントローラー
	ToolConditionsController := controllers.NewToolConditionsController()

	api.GET("/tool_conditions", ToolConditionsController.Index)
	// ロッドの硬さリスト取得
	api.GET("/tool_conditions/rod_hardness", ToolConditionsController.GetRodHardness)
	return e
}
