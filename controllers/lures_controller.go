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
func (uc *LuresController) GetAllLures(c echo.Context) error {
	db := database.GetDBConn()
	lures := []models.Lure{}
	db.Find(&lures)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		lures,
	))
}

/**
  ルアー取得
*/
func (uc *LuresController) GetLure(c echo.Context) error {
	db := database.GetDBConn()
	lure := models.Lure{}
	uid, _ := strconv.Atoi(c.Param("id"))
	db.First(&lure, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		lure,
	))
}

/**
  ルアー更新
*/
func (uc *LuresController) UpdateLure(c echo.Context) error {
	db := database.GetDBConn()
	lure := models.Lure{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&lure, uid)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	lure_type_id, _ := strconv.Atoi(c.FormValue("lure_type_id"))
	company_name := c.FormValue("company_name")
	color := c.FormValue("color")
	weight := c.FormValue("weight")
	deleteflg, _ := strconv.Atoi(c.FormValue("delete_flg"))

	db.Model(&lure).Updates(models.Lure{
		Name:        name,
		UserId:      user_id,
		LureTypeId:  lure_type_id,
		CompanyName: company_name,
		Color:       color,
		Weight:      weight,
		DeleteFlg:   deleteflg,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		lure,
	))
}

/**
  ルアー作成
*/
func (uc *LuresController) CreateLure(c echo.Context) error {
	db := database.GetDBConn()
	lure := models.Lure{}
	if err := c.Bind(&lure); err != nil {
		return err
	}

	result := db.Create(&lure).Error

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