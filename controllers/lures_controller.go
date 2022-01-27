package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// LuresController controller for Lures request
type LuresController struct{}

// NewLuresController is constructer for LuresController
func NewLuresController() *LuresController {
	return new(LuresController)
}

/**
  ルアー一覧取得
*/
func (uc *LuresController) Index(c echo.Context) error {
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	lures := []models.Lure{}
	result := models.GetAllLures(lures, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ルアー取得
*/
func (uc *LuresController) Show(c echo.Context) error {
	// idチェック
	lure_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	lure := models.Lure{}
	result := models.GetLure(lure, lure_id, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ルアー更新
*/
func (uc *LuresController) Update(c echo.Context) error {
	// idチェック
	lure_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データセット
	lure := models.Lure{}
	if err := c.Bind(&lure); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	lure.UserId = uid

	// 更新
	result := models.UpdateLure(lure, lure_id, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ルアー作成
*/
func (uc *LuresController) Create(c echo.Context) error {
	// データセット
	lure := models.Lure{}
	if err := c.Bind(&lure); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	lure.UserId = uid

	// 登録
	result := models.CreateLure(lure)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ルアー削除
*/
func (uc *LuresController) Delete(c echo.Context) error {
	lure := models.Lure{}
	lure_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.DeleteLure(lure, lure_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
