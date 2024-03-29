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
			weather.id as weather_id,
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
		GROUP BY weather_type, weather_id, color_name, color_code, result_type
		ORDER BY color_code ASC, result_type ASC, weather_id ASC
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
			depth.type_name as depth_type,
			depth.id as depth_id,
			CASE result.type_name
			  WHEN "no reaction" THEN "no reaction"
				ELSE "reaction"
			END as result_type
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
		GROUP BY depth_type, depth_id, color_name, color_code, result_type
		ORDER BY color_code ASC, result_type ASC, depth_id ASC
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
			lure_types.type_name as lure_type,
			lure_types.id as lure_type_id,
			CASE result.type_name
			  WHEN "no reaction" THEN "no reaction"
				ELSE "reaction"
			END as result_type
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
		GROUP BY color_name, color_code, result_type, lure_type, lure_type_id
		ORDER BY color_code ASC, result_type ASC, lure_type_id ASC
		`

	db.Raw(sql, uid, record_id, result_param).Scan(&result)
	return result
}
