-- 初期データ --
-- ユーザー --
INSERT INTO users 
(`id`, `name`, `email`, `password`, `nickname`, `first_name`, `last_name`, `group_id`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, "Yamada", "yamada@example.com", "testpass", "yama","taro", "yamada", 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, "Yoshida", "yoshida@example.com", "testpass", "yoshi", "go", "yoshida", 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, "Kawata", "kawata@example.com", "testpass", "kawa", "ken", "kawata", 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, "Hama", "hama@example.com", "testpass", "hana", "hanako", "hama", 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;

-- ルアータイプ --
INSERT INTO lure_types 
(`id`, `type_name`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, "spoon",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, "crank-bait", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, "minor", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, "viberation", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;

-- ヒットパターン各条件 --
INSERT INTO pattern_conditions 
(`id`, `type_num`, `type_name`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, 1, "caught",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, 1, "bit", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, 1, "chased", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, 2, "super fast",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(5, 2, "fast", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(6, 2, "normal", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(7, 2, "slow",  "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(8, 2, "super slow", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(9, 3, "top", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(10, 3, "shallow", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(11, 3, "middle", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(12, 3, "deep", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(13, 3, "bottom", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(14, 4, "sunny", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(15, 4, "rainy", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(16, 4, "cloudy", "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
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