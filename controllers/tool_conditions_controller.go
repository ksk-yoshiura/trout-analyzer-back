package controllers

import (
	"net/http"

	"trout-analyzer-back/models"

	"github.com/labstack/echo"
)

// ToolConditionsController controller for ToolConditions request
type ToolConditionsController struct{}

// NewToolConditionsController is constructer for ToolConditionsController
func NewToolConditionsController() *ToolConditionsController {
	return new(ToolConditionsController)
}

/**
  ツール条件各種一覧取得
*/
func (uc *ToolConditionsController) Index(c echo.Context) error {
	// データ取得
	tool_conditions := []models.ToolCondition{}
	result := models.GetAllToolConditions(tool_conditions)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
