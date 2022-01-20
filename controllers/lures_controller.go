package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
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
	lures := []models.Lure{}
	result := models.GetAllLures(lures)

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
	lure := models.Lure{}
	lure_id, _ := strconv.Atoi(c.Param("id"))
	result := models.GetLure(lure, lure_id)

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
	lure := models.Lure{}

	lure_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateLure(lure, lure_id, c)

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
	lure := models.Lure{}
	if err := c.Bind(&lure); err != nil {
		return err
	}

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
func (uc *LuresController) DeleteLure(c echo.Context) error {
	db := database.GetDBConn()
	lure := models.Lure{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&lure, uid)
	result := db.Delete(&lure)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
