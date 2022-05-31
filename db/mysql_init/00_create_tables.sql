-- CREATE DATABASE trout_analyzer;
use trout_analyzer;

SET GLOBAL sql_mode = 'ALLOW_INVALID_DATES';

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
  `nickname` VARCHAR(255) DEFAULT NULL,
  `first_name` VARCHAR(255) DEFAULT NULL,
  `last_name` VARCHAR(255) DEFAULT NULL,
  `group_id` tinyint(4) DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアーテーブル --

DROP TABLE IF EXISTS lures;

CREATE TABLE lures (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) DEFAULT NULL,
  `lure_type_id` tinyint(4) DEFAULT NULL,
  `user_id` int(255) DEFAULT NULL,
  `color_id` int(255) DEFAULT NULL,
  `sub_color_id` int(255) DEFAULT NULL,
  `weight` int(255) DEFAULT NULL,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアータイプテーブル --

DROP TABLE IF EXISTS lure_types;

CREATE TABLE lure_types (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ヒットパターン各条件テーブル --

DROP TABLE IF EXISTS pattern_conditions;

CREATE TABLE pattern_conditions (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_num`  int(255) DEFAULT NULL,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ツール各条件テーブル --

DROP TABLE IF EXISTS tool_conditions;

CREATE TABLE tool_conditions (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_num`  int(255) DEFAULT NULL,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- レコードテーブル --

DROP TABLE IF EXISTS records;

CREATE TABLE records (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `field_id` int(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ヒットパターンテーブル --

DROP TABLE IF EXISTS hit_patterns;

CREATE TABLE hit_patterns (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `lure_id` int(255) DEFAULT NULL,
  `tackle_id` int(255) DEFAULT NULL,
  `record_id` int(255) DEFAULT NULL,
  `speed` int(255) DEFAULT NULL,
  `depth` VARCHAR(255) DEFAULT NULL,
  `weather` VARCHAR(255) DEFAULT NULL,
  `result` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- タックルテーブル --

DROP TABLE IF EXISTS tackles;

CREATE TABLE tackles (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `rod_id` int(255) DEFAULT NULL,
  `reel_id` int(255) DEFAULT NULL,
  `line_id` int(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ロッドテーブル --

DROP TABLE IF EXISTS rods;

CREATE TABLE rods (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `hardness` tinyint(4) DEFAULT 0,
  `length` int(255) DEFAULT 0,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- リールテーブル --

DROP TABLE IF EXISTS reels;

CREATE TABLE reels (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `type_number` int(255) DEFAULT NULL,
  `gear` VARCHAR(255) DEFAULT NULL,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ラインテーブル --

DROP TABLE IF EXISTS fishing_lines;

CREATE TABLE fishing_lines (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `line_type_id` tinyint(4) DEFAULT 0,
  `thickness` tinyint(4) DEFAULT 0,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- フィールドテーブル --

DROP TABLE IF EXISTS fields;

CREATE TABLE fields (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `address` VARCHAR(255) DEFAULT NULL,
  `last_visited_at` timestamp DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- フィールド画像テーブル --

DROP TABLE IF EXISTS field_images;

CREATE TABLE field_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `field_id` int(255) DEFAULT NULL,
  `image_file` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ロッド画像テーブル --

DROP TABLE IF EXISTS rod_images;

CREATE TABLE rod_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `rod_id` int(255) DEFAULT NULL,
  `image_file` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- リール画像テーブル --

DROP TABLE IF EXISTS reel_images;

CREATE TABLE reel_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `reel_id` int(255) DEFAULT NULL,
  `image_file` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ライン画像テーブル --

DROP TABLE IF EXISTS line_images;

CREATE TABLE line_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `line_id` int(255) DEFAULT NULL,
  `image_file` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアー画像テーブル --

DROP TABLE IF EXISTS lure_images;

CREATE TABLE lure_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `lure_id` int(255) DEFAULT NULL,
  `image_file` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアーカラーテーブル --

DROP TABLE IF EXISTS colors;

CREATE TABLE colors (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) DEFAULT NULL,
  `code` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;