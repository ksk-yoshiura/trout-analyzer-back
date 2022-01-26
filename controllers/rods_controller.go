package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// RodsController controller for Rods request
type RodsController struct{}

// NewRodsController is constructer for RodsController
func NewRodsController() *RodsController {
	return new(RodsController)
}

/**
  ロッド一覧取得
*/
func (uc *RodsController) Index(c echo.Context) error {

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	rods := []models.Rod{}
	result := models.GetAllRods(rods, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ロッド取得
*/
func (uc *RodsController) Show(c echo.Context) error {
	// idチェック
	rod_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データ取得
	rod := models.Rod{}
	result := models.GetRod(rod, rod_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ロッド更新
*/
func (uc *RodsController) Update(c echo.Context) error {
	// idチェック
	rod_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データセット
	rod := models.Rod{}
	if err := c.Bind(&rod); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	rod.UserId = uid

	// 更新
	result := models.UpdateRod(rod, rod_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ロッド作成
*/
func (uc *RodsController) Create(c echo.Context) error {
	// データセット
	rod := models.Rod{}
	if err := c.Bind(&rod); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	rod.UserId = uid

	// 登録
	result := models.CreateRod(rod)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ロッド削除
*/
func (uc *RodsController) Delete(c echo.Context) error {
	// idチェック
	rod_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// 削除
	rod := models.Rod{}
	result := models.DeleteRod(rod, rod_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
