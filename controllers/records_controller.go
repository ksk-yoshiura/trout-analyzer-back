package controllers

import (
	"net/http"
	"strconv"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// RecordsController controller for Records request
type RecordsController struct{}

// NewRecordsController is constructer for RecordsController
func NewRecordsController() *RecordsController {
	return new(RecordsController)
}

/**
  レコード一覧取得
*/
func (uc *RecordsController) Index(c echo.Context) error {
	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	records := []models.Record{}
	result := models.GetAllRecords(records, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  レコード取得
*/
func (uc *RecordsController) Show(c echo.Context) error {
	// idチェック
	record_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	// データ取得
	record := models.Record{}
	result := models.GetRecord(record, record_id, uid)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  レコード更新
*/
func (uc *RecordsController) Update(c echo.Context) error {
	// idチェック
	record_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// データセット
	record := models.Record{}
	if err := c.Bind(&record); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	record.UserId = uid

	// 更新
	result := models.UpdateRecord(record, record_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  レコード作成
*/
func (uc *RecordsController) Create(c echo.Context) error {
	// データセット
	record := models.Record{}
	if err := c.Bind(&record); err != nil {
		return err
	}

	// トークンからユーザID取得
	uid := userIDFromToken(c)
	record.UserId = uid

	// 登録
	record, result := models.CreateRecord(record)

	// フィールド訪問最終日更新
	resultFieldUpdate := models.RecordLastVisitDate(uid, record.FieldId)

	if result == nil && resultFieldUpdate == nil {
		return c.JSON(http.StatusCreated, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			record,
		))
	} else {
		return c.JSON(http.StatusCreated, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			result,
		))
	}
}

/**
  レコード削除
*/
func (uc *RecordsController) Delete(c echo.Context) error {
	// idチェック
	record_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// 削除
	record := models.Record{}
	result := models.DeleteRecord(record, record_id)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}

/**
  レコード作成終了
*/
func (uc *RecordsController) Finish(c echo.Context) error {

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
	))
}
