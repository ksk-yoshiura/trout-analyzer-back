package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
)

// TacklesController controller for Tackles request
type TacklesController struct{}

// NewTacklesController is constructer for TacklesController
func NewTacklesController() *TacklesController {
	return new(TacklesController)
}

/**
  タックル一覧取得
*/
func (uc *TacklesController) Index(c echo.Context) error {
	tackles := []models.Tackle{}
	result := models.GetAllTackles(tackles)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  タックル取得
*/
func (uc *TacklesController) Show(c echo.Context) error {
	tackle := models.Tackle{}
	tackle_id, _ := strconv.Atoi(c.Param("id"))
	result := models.GetTackle(tackle, tackle_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  タックル更新
*/
func (uc *TacklesController) UpdateTackle(c echo.Context) error {
	db := database.GetDBConn()
	tackle := models.Tackle{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&tackle, uid)
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	rod_id, _ := strconv.Atoi(c.FormValue("rod_id"))
	reel_id, _ := strconv.Atoi(c.FormValue("reel_id"))
	line_id, _ := strconv.Atoi(c.FormValue("line_id"))

	db.Model(&tackle).Updates(models.Tackle{
		UserId: user_id,
		RodId:  rod_id,
		ReelId: reel_id,
		LineId: line_id,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		tackle,
	))
}

/**
  タックル作成
*/
func (uc *TacklesController) CreateTackle(c echo.Context) error {
	db := database.GetDBConn()
	tackle := models.Tackle{}
	if err := c.Bind(&tackle); err != nil {
		return err
	}

	result := db.Create(&tackle).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  タックル削除
*/
func (uc *TacklesController) DeleteTackle(c echo.Context) error {
	db := database.GetDBConn()
	tackle := models.Tackle{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&tackle, uid)
	result := db.Delete(&tackle)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
