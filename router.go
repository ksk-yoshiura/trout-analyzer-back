package main

import (
	"net/http"

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
	// CORS対策
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

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
	api.GET("/users", usersController.Show)
	api.PUT("/users", usersController.Update)
	api.POST("/users", usersController.Create)
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

	// レコードコントローラー
	recordsController := controllers.NewRecordsController()

	api.GET("/records/all", recordsController.Index)
	api.GET("/records/:id", recordsController.Show)
	api.PUT("/records/:id", recordsController.Update)
	api.POST("/records", recordsController.Create)
	api.POST("/records/:id", recordsController.Delete)

	// ヒットパターンコントローラー
	hitPatternsController := controllers.NewHitPatternsController()

	api.GET("/patterns", hitPatternsController.Index)
	api.GET("/patterns/:id", hitPatternsController.Show)
	api.PUT("/patterns/:id", hitPatternsController.Update)
	api.POST("/patterns", hitPatternsController.Create)
	api.POST("/patterns/:id", hitPatternsController.Delete)

	// ルアータイプコントローラー
	LureTypesController := controllers.NewLureTypesController()

	api.GET("/lure_types", LureTypesController.Index)

	// ツール条件各種コントローラー
	ToolConditionsController := controllers.NewToolConditionsController()

	api.GET("/tool_conditions", ToolConditionsController.Index)
	// ロッドの硬さリスト取得
	api.GET("/tool_conditions/rod_hardness", ToolConditionsController.GetRodHardness)
	// リールのギアリスト取得
	api.GET("/tool_conditions/reel_gear", ToolConditionsController.GetReelGear)
	// リールの型番リスト取得
	api.GET("/tool_conditions/reel_type", ToolConditionsController.GetReelType)
	// ラインタイプリスト取得
	api.GET("/tool_conditions/line_type", ToolConditionsController.GetLineType)

	// ヒットパターン条件各種コントローラー
	PatternConditionsController := controllers.NewPatternConditionsController()

	api.GET("/pattern_conditions", PatternConditionsController.Index)
	// 釣果条件各種一覧取得
	api.GET("/pattern_conditions/result", PatternConditionsController.GetResult)
	// ルアー速度条件各種一覧取得
	api.GET("/pattern_conditions/lure_speed", PatternConditionsController.GetLureSpeed)
	// ルアー深度条件各種一覧取得
	api.GET("/pattern_conditions/lure_depth", PatternConditionsController.GetLureDepth)
	// 天気条件各種一覧取得
	api.GET("/pattern_conditions/weather", PatternConditionsController.GetWeather)
	return e
}
