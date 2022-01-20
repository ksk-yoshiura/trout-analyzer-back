package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
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
	rods := []models.Rod{}
	result := models.GetAllRods(rods)

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
	rod := models.Rod{}
	rod_id, _ := strconv.Atoi(c.Param("id"))
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
	rod := models.Rod{}

	rod_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	result := models.UpdateRod(rod, rod_id, c)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ロッド作成
*/
func (uc *RodsController) CreateRod(c echo.Context) error {
	db := database.GetDBConn()
	rod := models.Rod{}
	if err := c.Bind(&rod); err != nil {
		return err
	}

	result := db.Create(&rod).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ロッド削除
*/
func (uc *RodsController) DeleteRod(c echo.Context) error {
	db := database.GetDBConn()
	rod := models.Rod{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&rod, uid)
	result := db.Delete(&rod)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
