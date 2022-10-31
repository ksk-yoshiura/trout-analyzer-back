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

-- 初期データ --

-- ルアータイプ --
INSERT INTO lure_types 
(`id`, `type_name`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, "spoon",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, "crank-bait", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, "minor", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, "viberation", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(5, "new type", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;

-- ヒットパターン各条件 --
INSERT INTO pattern_conditions 
(`id`, `type_num`, `type_name`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(null, 1, "caught",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 1, "bite", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 1, "chased", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 1, "no reaction", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 2, "super fast",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 2, "fast", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 2, "normal", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 2, "slow",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 2, "super slow", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 3, "top", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 3, "shallow", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 3, "middle", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 3, "deep", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 3, "bottom", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 4, "sunny", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 4, "rainy", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, 4, "cloudy", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;

-- ツール各条件 --
INSERT INTO tool_conditions
(`id`, `type_num`, `type_name`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, 1, "UL",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, 1, "L", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, 1, "ML", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, 1, "M", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(5, 2, "normal gear", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(6, 2, "hight gear", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(7, 3, "1000", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(8, 3, "1500", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(9, 3, "2000", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(10, 3, "2500", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(11, 4, "fluorocarbon", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(12, 4, "nylon", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(13, 4, "PE", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(14, 4, "ester", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;

-- ルアーカラー --
INSERT INTO colors
(`id`, `name`, `code`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(null, "black", "#000000",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "gray", "#808080", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "silver", "#C0C0C0", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "gold", "#FFD700", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "white", "#FFFFFF", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "beige", "#F5F5DC", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "blue", "#0000FF", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "navy", "#000080", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "teal", "#008080", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "green", "#008000", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "lime", "#00FF00", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "aqua", "#00FFFF", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "yellow", "#FFFF00", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "orange", "#FFA500", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "red", "#FF0000", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "fuchsia", "#FF00FF", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "olive", "#808000", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "purple", "#800080", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(null, "brown", "#A52A2A", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;




COMMIT;