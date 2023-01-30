package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo"
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
func (uc *ReelsController) Index(c echo.Context) error {
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	reels := []models.Reel{}
	result := models.GetAllReels(reels, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール取得
*/
func (uc *ReelsController) Show(c echo.Context) error {
	// idチェック
	reel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	reel := models.Reel{}
	result := models.GetReel(reel, reel_id, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール更新
*/
func (uc *ReelsController) Update(c echo.Context) error {
	// idチェック
	reel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データセット
	reel := models.Reel{}
	if err := c.Bind(&reel); err != nil {
		return err
	}

	// バリデーション
	if err := c.Validate(reel); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			c.Logger().Error(k + ": " + err.Error())
		}
		return err
	}

	// 画像
	image := models.Image{}
	if err := c.Bind(&image); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	reel.UserId = uid

	// 更新
	result := models.UpdateReel(reel, reel_id, image)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール作成
*/
func (uc *ReelsController) Create(c echo.Context) error {
	// データセット
	reel := models.Reel{}
	if err := c.Bind(&reel); err != nil {
		return err
	}

	// バリデーション
	if err := c.Validate(reel); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			c.Logger().Error(k + ": " + err.Error())
		}
		return err
	}

	// 画像
	image := models.Image{}
	if err := c.Bind(&image); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	reel.UserId = uid

	// 登録
	result := models.CreateReel(reel, image)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  リール削除
*/
func (uc *ReelsController) Delete(c echo.Context) error {
	// idチェック
	reel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// 削除
	reel := models.Reel{}
	result := models.DeleteReel(reel, reel_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
