
-- ユーザー --
INSERT INTO users 
(id, name, email,password, nickname, first_name, last_name, delete_flg, group_id, created_at, updated_at) 
VALUES 
(1, "Yamada", "yamada@example.com", "testpass", "yama","taro", "yamada", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00"),
(2, "Yoshida", "yoshida@example.com", "testpass", "yoshi", "go", "yoshida", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00"),
(3, "Kawata", "kawata@example.com", "testpass", "kawa", "ken", "kawata", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00"),
(4, "Hama", "hama@example.com", "testpass", "hana", "hanako", "hama", 0, 2, "2021-12-01 00:00:00", "2021-12-01 00:00:00")
;