package models

import (
	"trout-analyzer-back/database"

	"gorm.io/gorm"
)

type Color struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
}

/**
  カラー一覧取得
*/
func GetAllColors(colors []Color) []Color {
	db := database.GetDBConn()
	db.Find(&colors)
	return colors
}

/**
  カラー取得
*/
func GetColor(color Color, color_id int) Color {
	db := database.GetDBConn()
	db.First(&color, color_id)
	return color
}
