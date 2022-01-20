package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
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
	uid, _ := strconv.Atoi(c.Param("id"))
	result := models.GetField(field, uid)

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

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateField(field, uid, c)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  フィールド作成
*/
func (uc *FieldsController) CreateField(c echo.Context) error {
	db := database.GetDBConn()
	field := models.Field{}
	if err := c.Bind(&field); err != nil {
		return err
	}

	result := db.Create(&field).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  フィールド削除
*/
func (uc *FieldsController) DeleteField(c echo.Context) error {
	db := database.GetDBConn()
	field := models.Field{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&field, uid)
	result := db.Delete(&field)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
