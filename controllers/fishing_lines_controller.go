package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
)

// FishingLinesController controller for FishingLines request
type FishingLinesController struct{}

// NewFishingLinesController is constructer for FishingLinesController
func NewFishingLinesController() *FishingLinesController {
	return new(FishingLinesController)
}

/**
  ライン一覧取得
*/
func (uc *FishingLinesController) GetAllFishingLines(c echo.Context) error {
	db := database.GetDBConn()
	fishing_lines := []models.FishingLine{}
	db.Find(&fishing_lines)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		fishing_lines,
	))
}

/**
  ライン取得
*/
func (uc *FishingLinesController) GetFishingLine(c echo.Context) error {
	db := database.GetDBConn()
	fishing_line := models.FishingLine{}
	uid, _ := strconv.Atoi(c.Param("id"))
	db.First(&fishing_line, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		fishing_line,
	))
}

/**
  ライン更新
*/
func (uc *FishingLinesController) UpdateFishingLine(c echo.Context) error {
	db := database.GetDBConn()
	fishing_line := models.FishingLine{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&fishing_line, uid)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	line_type_id, _ := strconv.Atoi(c.FormValue("line_type_id"))
	thickness, _ := strconv.Atoi(c.FormValue("thickness"))
	company_name := c.FormValue("company_name")

	db.Model(&fishing_line).Updates(models.FishingLine{
		Name:        name,
		UserId:      user_id,
		LineTypeId:  line_type_id,
		Thickness:   thickness,
		CompanyName: company_name,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		fishing_line,
	))
}

/**
  ライン作成
*/
func (uc *FishingLinesController) CreateFishingLine(c echo.Context) error {
	db := database.GetDBConn()
	fishing_line := models.FishingLine{}
	if err := c.Bind(&fishing_line); err != nil {
		return err
	}

	result := db.Create(&fishing_line).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  ライン削除
*/
func (uc *FishingLinesController) DeleteFishingLine(c echo.Context) error {
	db := database.GetDBConn()
	fishing_line := models.FishingLine{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&fishing_line, uid)
	result := db.Delete(&fishing_line)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
