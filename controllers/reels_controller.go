package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// ReelsController controller for Reels request
type ReelsController struct{}

// NewReelsController is constructer for ReelsController
func NewReelsController() *ReelsController {
	return new(ReelsController)
}

/**
  リール一覧取得
*/
func (uc *ReelsController) Index(c echo.Context) error {
	reels := []models.Reel{}
	result := models.GetAllReels(reels)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール取得
*/
func (uc *ReelsController) Show(c echo.Context) error {
	// idチェック
	reel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データ取得
	reel := models.Reel{}
	result := models.GetReel(reel, reel_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール更新
*/
func (uc *ReelsController) Update(c echo.Context) error {
	reel := models.Reel{}
	reel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateReel(reel, reel_id, c)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール作成
*/
func (uc *ReelsController) Create(c echo.Context) error {
	reel := models.Reel{}
	if err := c.Bind(&reel); err != nil {
		return err
	}

	result := models.CreateReel(reel)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール削除
*/
func (uc *ReelsController) Delete(c echo.Context) error {
	reel := models.Reel{}
	reel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.DeleteReel(reel, reel_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
