package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
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
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	fishing_lines := []models.FishingLine{}
	result := models.GetAllLines(fishing_lines, uid)

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
	// idチェック
	line_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	fishing_line := models.FishingLine{}
	result := models.GetLine(fishing_line, line_id, uid)

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
	// idチェック
	line_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データセット
	fishing_line := models.FishingLine{}
	if err := c.Bind(&fishing_line); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	fishing_line.UserId = uid

	result := models.UpdateLine(fishing_line, line_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン作成
*/
func (uc *FishingLinesController) Create(c echo.Context) error {
	// データセット
	fishing_line := models.FishingLine{}
	if err := c.Bind(&fishing_line); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	fishing_line.UserId = uid

	// 登録
	result := models.CreateLine(fishing_line)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン削除
*/
func (uc *FishingLinesController) Delete(c echo.Context) error {
	// idチェック
	line_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// 削除
	fishing_line := models.FishingLine{}
	result := models.DeleteLine(fishing_line, line_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
