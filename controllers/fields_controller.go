package controllers

import (
	"net/http"
	"strconv"

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
	// データ取得
	fields := []models.Field{}
	result := models.GetAllFields(fields)

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
	field := models.Field{}
	field_id, _ := strconv.Atoi(c.Param("id"))
	result := models.GetField(field, field_id)

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
	field := models.Field{}

	field_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateField(field, field_id, c)

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
	field := models.Field{}
	if err := c.Bind(&field); err != nil {
		return err
	}

	result := models.CreateField(field)

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
