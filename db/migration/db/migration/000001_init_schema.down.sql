DROP TABLE IF EXISTS users;

-- ルアーテーブル --

DROP TABLE IF EXISTS lures;

-- ルアータイプテーブル --

DROP TABLE IF EXISTS lure_types;

-- ヒットパターン各条件テーブル --

DROP TABLE IF EXISTS pattern_conditions;

-- ツール各条件テーブル --

DROP TABLE IF EXISTS tool_conditions;

-- レコードテーブル --

DROP TABLE IF EXISTS records;

-- ヒットパターンテーブル --

DROP TABLE IF EXISTS hit_patterns;

-- タックルテーブル --

DROP TABLE IF EXISTS tackles;

CREATE TABLE tackles (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `rod_id` int(255) NOT NULL,
  `reel_id` int(255) NOT NULL,
  `line_id` int(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ロッドテーブル --

DROP TABLE IF EXISTS rods;

-- リールテーブル --

DROP TABLE IF EXISTS reels;

-- ラインテーブル --

DROP TABLE IF EXISTS fishing_lines;

-- フィールドテーブル --

DROP TABLE IF EXISTS fields;

-- フィールド画像テーブル --

DROP TABLE IF EXISTS field_images;

-- ロッド画像テーブル --

DROP TABLE IF EXISTS rod_images;

-- リール画像テーブル --

DROP TABLE IF EXISTS reel_images;

-- ライン画像テーブル --

DROP TABLE IF EXISTS line_images;

-- ルアー画像テーブル --

DROP TABLE IF EXISTS lure_images;

-- ルアーカラーテーブル --

DROP TABLE IF EXISTS colors;