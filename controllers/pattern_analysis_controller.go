package controllers

import (
	"net/http"
	"strconv"
	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// PatternAnalysisController controller for HitPatterns request
type PatternAnalysisController struct{}

// NewPatternAnalysisController is constructer for PatternAnalysisController
func NewPatternAnalysisController() *PatternAnalysisController {
	return new(PatternAnalysisController)
}

/**
  ヒットパターン分析結果取得
	ルアーカラーと天気
*/
func (uc *PatternAnalysisController) GetColorWeatherRelation(c echo.Context) error {
	// idチェック
	record_id, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		return echo.ErrNotFound
	}
	// resultパラメータチェック
	result_param := c.Param("result")

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	result := models.GetColorWeatherAnalysis(result_param, uid, record_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン分析結果取得
	ルアーカラーと深度
*/
func (uc *PatternAnalysisController) GetColorDepthRelation(c echo.Context) error {
	// idチェック
	record_id, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		return echo.ErrNotFound
	}
	// resultパラメータチェック
	result_param := c.Param("result")

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	result := models.GetColorDepthAnalysis(result_param, uid, record_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ヒットパターン分析結果取得
	ルアーカラーとルアータイプ
*/
func (uc *PatternAnalysisController) GetColorLureTypeRelation(c echo.Context) error {

	// // トークンからユーザID取得
	// uid := userIDFromToken(c)
	// // データ取得
	// hit_pattern := []models.HitPattern{}
	// result := models.GetAnalysis(hit_pattern, uid, record_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"color type",
	))
}
