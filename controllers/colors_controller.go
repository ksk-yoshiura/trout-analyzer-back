package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// ColorsController controller for Colors request
type ColorsController struct{}

// NewColorsController is constructer for ColorsController
func NewColorsController() *ColorsController {
	return new(ColorsController)
}

/**
  カラー一覧取得
*/
func (uc *ColorsController) Index(c echo.Context) error {
	// データ取得
	colors := []models.Color{}
	result := models.GetAllColors(colors)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  カラー取得
*/
func (uc *ColorsController) Show(c echo.Context) error {
	// idチェック
	color_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データ取得
	color := models.Color{}
	result := models.GetColor(color, color_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
