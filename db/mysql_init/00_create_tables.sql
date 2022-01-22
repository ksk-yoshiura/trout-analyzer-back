-- CREATE DATABASE trout_analyzer;
use trout_analyzer;


DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
  `nickname` VARCHAR(255) DEFAULT NULL,
  `first_name` VARCHAR(255) DEFAULT NULL,
  `last_name` VARCHAR(255) DEFAULT NULL,
  `group_id` tinyint(4) DEFAULT 1,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアーテーブル --

DROP TABLE IF EXISTS lures;

CREATE TABLE lures (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) DEFAULT NULL,
  `lure_type_id` tinyint(4) DEFAULT NULL,
  `user_id` int(255) DEFAULT NULL,
  `color` VARCHAR(255) DEFAULT NULL,
  `weight` int(255) DEFAULT NULL,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアータイプテーブル --

DROP TABLE IF EXISTS lure_types;

CREATE TABLE lure_types (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ヒットパターン各条件テーブル --

DROP TABLE IF EXISTS pattern_conditions;

CREATE TABLE pattern_conditions (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_num`  int(255) DEFAULT NULL,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ツール各条件テーブル --

DROP TABLE IF EXISTS tool_conditions;

CREATE TABLE tool_conditions (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_num`  int(255) DEFAULT NULL,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ヒットパターンテーブル --

DROP TABLE IF EXISTS hit_patterns;

CREATE TABLE hit_patterns (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `lure_id` int(255) DEFAULT NULL,
  `tackle_id` int(255) DEFAULT NULL,
  `speed` int(255) DEFAULT NULL,
  `depth` VARCHAR(255) DEFAULT NULL,
  `weather` VARCHAR(255) DEFAULT NULL,
  `result` VARCHAR(255) DEFAULT NULL,
  `field_id` int(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
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
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ロッドテーブル --

DROP TABLE IF EXISTS rods;

CREATE TABLE rods (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `hardness_id` tinyint(4) DEFAULT 0,
  `length` int(255) DEFAULT 0,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
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
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
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
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- フィールドテーブル --

DROP TABLE IF EXISTS fields;

CREATE TABLE fields (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) DEFAULT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `address` VARCHAR(255) DEFAULT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  `deleted_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;