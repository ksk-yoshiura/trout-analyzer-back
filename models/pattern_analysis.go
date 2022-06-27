package models

import (
	"trout-analyzer-back/database"
)

type ColorWeatherResult struct {
	Sum         int
	ColorName   string
	ColorCode   string
	ResultType  string
	WeatherType string
}

type ColorDepthResult struct {
	Sum        int
	ColorName  string
	ColorCode  string
	ResultType string
	DepthType  string
}

type ColorLureTypehResult struct {
	Sum        int
	ColorName  string
	ColorCode  string
	ResultType string
	LureType   string
}

const pattern_weather = 4
const pattern_depth = 3

/**
  ヒットパッテーン分析取得
	ルアーカラーと天気
*/
func GetColorWeatherAnalysis(result_param string, uid int, record_id int) []ColorWeatherResult {
	var result []ColorWeatherResult
	db := database.GetDBConn()
	// ログインユーザは自分の分析しか見れない
	sql :=
		`SELECT 
			COUNT(colors.id) as sum, 
			colors.name as color_name, 
			colors.code as color_code,
			weather.type_name as weather_type,
			CASE result.type_name
			  WHEN "no reaction" THEN "no reaction"
				ELSE "reaction"
			END as result_type
		FROM hit_patterns 
		LEFT JOIN lures ON lures.id = hit_patterns.lure_id
		LEFT JOIN colors ON lures.color_id = colors.id
		LEFT JOIN pattern_conditions as weather ON weather.id = hit_patterns.weather
		LEFT JOIN pattern_conditions as result ON result.id = hit_patterns.result
		WHERE hit_patterns.user_id = ? AND hit_patterns.record_id = ? AND weather.type_num = ?
		`
	if result_param != "all" {
		sql += `AND result.type_name = ? OR result.type_name = "no reaction"`
	} else { // FIX ME：一時的な誤魔化し
		sql += `AND result.type_name != ?`
	}
	sql +=
		`
		GROUP BY color_name, color_code, result_type, weather_type
		`

	db.Raw(sql, uid, record_id, pattern_weather, result_param).Scan(&result)
	return result
}

/**
  ヒットパッテーン分析取得
	ルアーカラーと深度
*/
func GetColorDepthAnalysis(result_param string, uid int, record_id int) []ColorDepthResult {
	var result []ColorDepthResult
	db := database.GetDBConn()
	// ログインユーザは自分の分析しか見れない
	sql :=
		`SELECT 
			COUNT(colors.id) as sum, 
			colors.name as color_name, 
			colors.code as color_code,
			result.type_name as result_type,
			depth.type_name as depth_type
		FROM hit_patterns 
		LEFT JOIN lures ON lures.id = hit_patterns.lure_id
		LEFT JOIN colors ON lures.color_id = colors.id
		LEFT JOIN pattern_conditions as depth ON depth.id = hit_patterns.depth
		LEFT JOIN pattern_conditions as result ON result.id = hit_patterns.result
		WHERE hit_patterns.user_id = ? AND hit_patterns.record_id = ? AND depth.type_num = ?
		`
	if result_param != "all" {
		sql += `AND result.type_name = ? OR result.type_name = "no reaction"`
	} else { // FIX ME：一時的な誤魔化し
		sql += `AND result.type_name != ?`
	}
	sql +=
		`
		GROUP BY color_name, color_code, result_type, depth_type
		`

	db.Raw(sql, uid, record_id, pattern_depth, result_param).Scan(&result)
	return result
}

/**
  ヒットパッテーン分析取得
	ルアーカラーとルアータイプ
*/
func GetColorLureTypeAnalysis(result_param string, uid int, record_id int) []ColorLureTypehResult {
	var result []ColorLureTypehResult
	db := database.GetDBConn()
	// ログインユーザは自分の分析しか見れない
	sql :=
		`SELECT 
			COUNT(colors.id) as sum, 
			colors.name as color_name, 
			colors.code as color_code,
			result.type_name as result_type,
			lure_types.type_name as lure_type
		FROM hit_patterns 
		LEFT JOIN lures ON lures.id = hit_patterns.lure_id
		LEFT JOIN lure_types ON lure_types.id = lures.lure_type_id
		LEFT JOIN colors ON lures.color_id = colors.id
		LEFT JOIN pattern_conditions as result ON result.id = hit_patterns.result
		WHERE hit_patterns.user_id = ? AND hit_patterns.record_id = ?
		`
	if result_param != "all" {
		sql += `AND result.type_name = ? OR result.type_name = "no reaction"`
	} else { // FIX ME：一時的な誤魔化し
		sql += `AND result.type_name != ?`
	}
	sql +=
		`
		GROUP BY color_name, color_code, result_type, lure_type
		`

	db.Raw(sql, uid, record_id, result_param).Scan(&result)
	return result
}
