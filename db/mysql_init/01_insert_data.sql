-- 初期データ --
-- ユーザー --
INSERT INTO users 
(`id`, `name`, `email`, `password`, `nickname`, `first_name`, `last_name`, `delete_flg`, `group_id`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, "Yamada", "yamada@example.com", "testpass", "yama","taro", "yamada", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, "Yoshida", "yoshida@example.com", "testpass", "yoshi", "go", "yoshida", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, "Kawata", "kawata@example.com", "testpass", "kawa", "ken", "kawata", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, "Hama", "hama@example.com", "testpass", "hana", "hanako", "hama", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;

-- ルアータイプ --
INSERT INTO lure_types 
(`id`, `type_name`, `delete_flg`, `created_at`, `updated_at`, `deleted_at`) 
VALUES 
(1, "spoon",  0, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(2, "crank-bait", 0, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(3, "minor", 0, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null),
(4, "viberation", 0, "2021-12-01 00:00:00", "2021-12-01 00:00:00", null)
;