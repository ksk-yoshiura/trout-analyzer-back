package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo"
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
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	tackles := []models.Tackle{}
	result := models.GetAllTackles(tackles, uid)

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
	// idチェック
	tackle_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	tackle := models.Tackle{}
	result := models.GetTackle(tackle, tackle_id, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  タックル更新
*/
func (uc *TacklesController) Update(c echo.Context) error {
	tackle_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	tackle := models.Tackle{}
	if err := c.Bind(&tackle); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	tackle.UserId = uid

	// 更新
	result := models.UpdateTackle(tackle, tackle_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  タックル作成
*/
func (uc *TacklesController) Create(c echo.Context) error {
	// データセット
	tackle := models.Tackle{}
	if err := c.Bind(&tackle); err != nil {
		return err
	}

	// バリデーション
	if err := c.Validate(tackle); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			c.Logger().Error(k + ": " + err.Error())
		}
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	tackle.UserId = uid

	// 登録
	result := models.CreateTackle(tackle)

	return c.JSON(http.StatusCreated, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  タックル削除
*/
func (uc *TacklesController) Delete(c echo.Context) error {
	// idチェック
	tackle_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// 削除
	tackle := models.Tackle{}
	result := models.DeleteTackle(tackle, tackle_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
