package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"

	"trout-analyzer-back/database"
)

// ReelsController controller for Reels request
type ReelsController struct{}

// NewReelsController is constructer for ReelsController
func NewReelsController() *ReelsController {
	return new(ReelsController)
}

/**
  リール一覧取得
*/
func (uc *ReelsController) GetAllReels(c echo.Context) error {
	db := database.GetDBConn()
	Reels := []models.Reel{}
	db.Find(&Reels)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		Reels,
	))
}

/**
  リール取得
*/
func (uc *ReelsController) GetReel(c echo.Context) error {
	db := database.GetDBConn()
	reel := models.Reel{}
	uid, _ := strconv.Atoi(c.Param("id"))
	db.First(&reel, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		reel,
	))
}

/**
  リール更新
*/
func (uc *ReelsController) UpdateReel(c echo.Context) error {
	db := database.GetDBConn()
	reel := models.Reel{}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&reel, uid)
	name := c.FormValue("name")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	type_number := c.FormValue("type_number")
	gear := c.FormValue("gear")
	company_name := c.FormValue("company_name")
	deleteflg, _ := strconv.Atoi(c.FormValue("delete_flg"))

	db.Model(&reel).Updates(models.Reel{
		Name:        name,
		UserId:      user_id,
		TypeNumber:  type_number,
		Gear:        gear,
		CompanyName: company_name,
		DeleteFlg:   deleteflg,
	})

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		reel,
	))
}

/**
  リール作成
*/
func (uc *ReelsController) CreateReel(c echo.Context) error {
	db := database.GetDBConn()
	reel := models.Reel{}
	if err := c.Bind(&reel); err != nil {
		return err
	}

	result := db.Create(&reel).Error

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール削除
*/
func (uc *ReelsController) DeleteReel(c echo.Context) error {
	db := database.GetDBConn()
	reel := models.Reel{}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	db.First(&reel, uid)
	result := db.Delete(&reel)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
