package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
)

// HitPatternsController controller for HitPatterns request
type HitPatternsController struct{}

// NewHitPatternsController is constructer for HitPatternsController
func NewHitPatternsController() *HitPatternsController {
	return new(HitPatternsController)
}

/**
  ヒットパターン一覧取得
*/
func (uc *HitPatternsController) Index(c echo.Context) error {
	hit_patterns := []models.HitPattern{}
	result := models.GetAllHitPatterns(hit_patterns)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン取得
*/
func (uc *HitPatternsController) Show(c echo.Context) error {
	hit_pattern := models.HitPattern{}
	hit_pattern_id, _ := strconv.Atoi(c.Param("id"))
	result := models.GetHitPattern(hit_pattern, hit_pattern_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン更新
*/
func (uc *HitPatternsController) UpdateHitPattern(c echo.Context) error {
	db := database.GetDBConn()
	hit_pattern := models.HitPattern{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&hit_pattern, uid)
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	lure_id, _ := strconv.Atoi(c.FormValue("lure_id"))
	tackle_id, _ := strconv.Atoi(c.FormValue("tackle_id"))
	speed, _ := strconv.Atoi(c.FormValue("speed"))
	depth, _ := strconv.Atoi(c.FormValue("depth"))
	weather, _ := strconv.Atoi(c.FormValue("weather"))
	result, _ := strconv.Atoi(c.FormValue("result"))
	field_id, _ := strconv.Atoi(c.FormValue("field_id"))

	db.Model(&hit_pattern).Updates(models.HitPattern{
		UserId:   user_id,
		LureId:   lure_id,
		TackleId: tackle_id,
		Speed:    speed,
		Depth:    depth,
		Weather:  weather,
		Result:   result,
		FieldId:  field_id,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		hit_pattern,
	))
}

/**
  ヒットパターン作成
*/
func (uc *HitPatternsController) CreateHitPattern(c echo.Context) error {
	db := database.GetDBConn()
	hit_pattern := models.HitPattern{}
	if err := c.Bind(&hit_pattern); err != nil {
		return err
	}

	result := db.Create(&hit_pattern).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン削除
*/
func (uc *HitPatternsController) DeleteHitPattern(c echo.Context) error {
	db := database.GetDBConn()
	hit_pattern := models.HitPattern{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&hit_pattern, uid)
	result := db.Delete(&hit_pattern)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
