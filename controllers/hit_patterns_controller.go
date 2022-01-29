package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
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
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	hit_patterns := []models.HitPattern{}
	result := models.GetAllHitPatterns(hit_patterns, uid)

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
func (uc *HitPatternsController) Update(c echo.Context) error {
	hit_pattern := models.HitPattern{}

	hit_pattern_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateHitPattern(hit_pattern, hit_pattern_id, c)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン作成
*/
func (uc *HitPatternsController) Create(c echo.Context) error {
	hit_pattern := models.HitPattern{}
	if err := c.Bind(&hit_pattern); err != nil {
		return err
	}

	result := models.CreateHitPattern(hit_pattern)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン削除
*/
func (uc *HitPatternsController) Delete(c echo.Context) error {
	hit_pattern := models.HitPattern{}
	hit_pattern_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.DeleteHitPattern(hit_pattern, hit_pattern_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
