package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
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
	reel := models.Reel{}
	reel_id, _ := strconv.Atoi(c.Param("id"))
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
func (uc *ReelsController) CreateReel(c echo.Context) error {
	db := database.GetDBConn()
	reel := models.Reel{}
	if err := c.Bind(&reel); err != nil {
		return err
	}

	result := db.Create(&reel).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール削除
*/
func (uc *ReelsController) DeleteReel(c echo.Context) error {
	db := database.GetDBConn()
	reel := models.Reel{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&reel, uid)
	result := db.Delete(&reel)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
