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
	// クエリパラメータ
	// ルアータイプ
	type_id := c.QueryParam("type_id")

	var result []models.Lure
	// ルアータイプIDがある場合絞り込む
	if type_id == "0" || len(type_id) == 0 {
		result = models.GetAllLures(lures, uid)
	} else {
		result = models.GetLuresSelectedLureType(lures, type_id, uid)

	}

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
	// 画像
	image := models.Image{}
	if err := c.Bind(&image); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	lure.UserId = uid

	// 登録
	result := models.CreateLure(lure, image)

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
	// idチェック
	lure_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// 削除
	lure := models.Lure{}
	result := models.DeleteLure(lure, lure_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
