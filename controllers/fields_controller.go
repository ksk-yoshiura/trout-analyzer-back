package controllers

import (
	"net/http"
	"strconv"
	"time"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// FieldsController controller for Fields request
type FieldsController struct{}

// NewFieldsController is constructer for FieldsController
func NewFieldsController() *FieldsController {
	return new(FieldsController)
}

/**
  フィールド一覧取得
*/
func (uc *FieldsController) Index(c echo.Context) error {
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	fields := []models.Field{}
	result := models.GetAllFields(fields, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  フィールド取得
*/
func (uc *FieldsController) Show(c echo.Context) error {
	// idチェック
	field_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	field := models.Field{}
	result := models.GetField(field, field_id, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  フィールド更新
*/
func (uc *FieldsController) Update(c echo.Context) error {
	// idチェック
	field_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データセット
	field := models.Field{}
	if err := c.Bind(&field); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	field.UserId = uid

	// 更新
	result := models.UpdateField(field, field_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  フィールド作成
*/
func (uc *FieldsController) Create(c echo.Context) error {
	// データセット
	field := models.Field{}
	if err := c.Bind(&field); err != nil {
		return err
	}
	// 画像
	image := models.Image{}
	if err := c.Bind(&image); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	field.UserId = uid
	// 新規登録の時はからデータを入れる
	field.LastVisitedAt = time.Time{}

	// 登録
	result := models.CreateField(field, image)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  フィールド削除
*/
func (uc *FieldsController) Delete(c echo.Context) error {
	field := models.Field{}
	field_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.DeleteField(field, field_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
