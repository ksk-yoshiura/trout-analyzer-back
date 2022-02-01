package controllers

import (
	"net/http"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// PatternConditionsController controller for PatternConditions request
type PatternConditionsController struct{}

// NewPatternConditionsController is constructer for PatternConditionsController
func NewPatternConditionsController() *PatternConditionsController {
	return new(PatternConditionsController)
}

/**
  ヒットパターン条件各種一覧取得
*/
func (uc *PatternConditionsController) Index(c echo.Context) error {
	// データ取得
	pattern_conditions := []models.PatternCondition{}
	result := models.GetAllPatternConditions(pattern_conditions)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  釣果条件各種一覧取得
*/
func (uc *PatternConditionsController) GetResult(c echo.Context) error {
	// データ取得
	pattern_conditions := []models.PatternCondition{}
	result := models.GetResultConditions(pattern_conditions)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ルアー速度条件各種一覧取得
*/
func (uc *PatternConditionsController) GetLureSpeed(c echo.Context) error {
	// データ取得
	pattern_conditions := []models.PatternCondition{}
	result := models.GetLureSpeedConditions(pattern_conditions)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ルアー深度条件各種一覧取得
*/
func (uc *PatternConditionsController) GetLureDepth(c echo.Context) error {
	// データ取得
	pattern_conditions := []models.PatternCondition{}
	result := models.GetLureDepthConditions(pattern_conditions)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  天気条件各種一覧取得
*/
func (uc *PatternConditionsController) GetWeather(c echo.Context) error {
	// データ取得
	pattern_conditions := []models.PatternCondition{}
	result := models.GetWeatherConditions(pattern_conditions)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
