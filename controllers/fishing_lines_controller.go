package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
)

// FishingLinesController controller for FishingLines request
type FishingLinesController struct{}

// NewFishingLinesController is constructer for FishingLinesController
func NewFishingLinesController() *FishingLinesController {
	return new(FishingLinesController)
}

/**
  ライン一覧取得
*/
func (uc *FishingLinesController) Index(c echo.Context) error {
	fishing_lines := []models.FishingLine{}
	result := models.GetAllLines(fishing_lines)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン取得
*/
func (uc *FishingLinesController) Show(c echo.Context) error {
	fishing_line := models.FishingLine{}
	line_id, _ := strconv.Atoi(c.Param("id"))
	result := models.GetLine(fishing_line, line_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン更新
*/
func (uc *FishingLinesController) Update(c echo.Context) error {
	fishing_line := models.FishingLine{}
	line_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateLine(fishing_line, line_id, c)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン作成
*/
func (uc *FishingLinesController) CreateFishingLine(c echo.Context) error {
	db := database.GetDBConn()
	fishing_line := models.FishingLine{}
	if err := c.Bind(&fishing_line); err != nil {
		return err
	}

	result := db.Create(&fishing_line).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン削除
*/
func (uc *FishingLinesController) DeleteFishingLine(c echo.Context) error {
	db := database.GetDBConn()
	fishing_line := models.FishingLine{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&fishing_line, uid)
	result := db.Delete(&fishing_line)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
