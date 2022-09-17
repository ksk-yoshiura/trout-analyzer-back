BEGIN;
-- ユーザーテーブル --

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

CREATE TABLE lures (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) DEFAULT NULL,
  `lure_type_id` tinyint(4) NOT NULL,
  `user_id` int(255) NOT NULL,
  `color_id` int(255) NOT NULL,
  `sub_color_id` int(255) DEFAULT NULL,
  `weight` int(255) DEFAULT NULL,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアータイプテーブル --

CREATE TABLE lure_types (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ヒットパターン各条件テーブル --

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

CREATE TABLE records (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `field_id` int(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ヒットパターンテーブル --

CREATE TABLE hit_patterns (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `lure_id` int(255) DEFAULT NULL,
  `tackle_id` int(255) NOT NULL,
  `record_id` int(255) NOT NULL,
  `speed` int(255) NOT NULL,
  `depth` VARCHAR(255) NOT NULL,
  `weather` VARCHAR(255) NOT NULL,
  `result` VARCHAR(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- タックルテーブル --

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

CREATE TABLE rods (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `hardness` tinyint(4) DEFAULT 0,
  `length` int(255) DEFAULT 0,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- リールテーブル --

CREATE TABLE reels (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `type_number` int(255) DEFAULT NULL,
  `gear` VARCHAR(255) DEFAULT NULL,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ラインテーブル --

CREATE TABLE fishing_lines (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `line_type_id` tinyint(4) DEFAULT 0,
  `thickness` tinyint(4) DEFAULT 0,
  `company_name` VARCHAR(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- フィールドテーブル --

CREATE TABLE fields (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` int(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) DEFAULT NULL,
  `last_visited_at` timestamp DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- フィールド画像テーブル --

CREATE TABLE field_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `field_id` int(255) NOT NULL,
  `image_file` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ロッド画像テーブル --

CREATE TABLE rod_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `rod_id` int(255) NOT NULL,
  `image_file` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- リール画像テーブル --

CREATE TABLE reel_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `reel_id` int(255) NOT NULL,
  `image_file` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ライン画像テーブル --

CREATE TABLE line_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `line_id` int(255) NOT NULL,
  `image_file` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアー画像テーブル --

CREATE TABLE lure_images (
  `id` INT NOT NULL AUTO_INCREMENT,
  `lure_id` int(255) NOT NULL,
  `image_file` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ルアーカラーテーブル --

CREATE TABLE colors (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) DEFAULT NULL,
  `code` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

COMMIT;