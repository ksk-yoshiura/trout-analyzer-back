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
func (uc *RodsController) GetRod(c echo.Context) error {
	db := database.GetDBConn()
	rod := models.Rod{}
	uid, _ := strconv.Atoi(c.Param("id"))
	db.First(&rod, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		rod,
	))
}

/**
  ロッド更新
*/
func (uc *RodsController) UpdateRod(c echo.Context) error {
	db := database.GetDBConn()
	rod := models.Rod{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&rod, uid)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	hardness_id, _ := strconv.Atoi(c.FormValue("hardness_id"))
	length := c.FormValue("length")
	company_name := c.FormValue("company_name")

	db.Model(&rod).Updates(models.Rod{
		Name:        name,
		UserId:      user_id,
		HardnessId:  hardness_id,
		Length:      length,
		CompanyName: company_name,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		rod,
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
