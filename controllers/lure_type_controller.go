package controllers

import (
	"net/http"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// LureTypesController controller for LureTypes request
type LureTypesController struct{}

// NewLureTypesController is constructer for LureTypesController
func NewLureTypesController() *LureTypesController {
	return new(LureTypesController)
}

/**
  ルアータイプ一覧取得
*/
func (uc *LureTypesController) Index(c echo.Context) error {
	// データ取得
	lure_types := []models.LureType{}
	result := models.GetAllLureTypes(lure_types)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
