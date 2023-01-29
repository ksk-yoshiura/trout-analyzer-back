BEGIN;

-- ユーザーテーブル
-- サインアップ時
ALTER TABLE users ADD INDEX email_index(email);

-- ルアーテーブル
-- ルアー一覧
ALTER TABLE lures ADD INDEX user_id_index(user_id);

-- ルアー画像テーブル
-- ルアー一覧
ALTER TABLE lure_images ADD INDEX lure_id_index(lure_id);

-- ロッドテーブル
-- ロッド一覧
ALTER TABLE rods ADD INDEX user_id_index(user_id);

-- ロッド画像テーブル
-- ロッド一覧
ALTER TABLE rod_images ADD INDEX rod_id_index(rod_id);

-- リールテーブル
-- リール一覧
ALTER TABLE reels ADD INDEX user_id_index(user_id);

-- リール画像テーブル
-- リール一覧
ALTER TABLE reel_images ADD INDEX reel_id_index(reel_id);

-- ラインテーブル
-- ライン一覧
ALTER TABLE fishing_lines ADD INDEX user_id_index(user_id);

-- ライン画像テーブル
-- ライン一覧
ALTER TABLE line_images ADD INDEX line_id_index(line_id);

-- フィールドテーブル
-- フィールド一覧
ALTER TABLE fields ADD INDEX user_id_index(user_id);

-- フィールド画像テーブル
-- フィールド一覧
ALTER TABLE field_images ADD INDEX field_id_index(field_id);

-- タックルテーブル
-- タックル一覧
ALTER TABLE tackles ADD INDEX user_id_index(user_id);

-- レコードテーブル
-- レコード一覧
ALTER TABLE records ADD INDEX user_id_index(user_id);

-- ヒットパターンテーブル
-- ヒットパターン一覧
ALTER TABLE hit_patterns ADD INDEX user_id_index(user_id);
ALTER TABLE hit_patterns ADD INDEX record_id_index(record_id);

-- パターン条件テーブル
-- パターン条件一覧
ALTER TABLE pattern_conditions ADD INDEX type_num_index(type_num);

-- ルアータイプテーブル
-- ルアータイプ一覧
ALTER TABLE lure_types ADD INDEX type_name_index(type_name);

-- カラーテーブル
-- カラー一覧
ALTER TABLE colors MODIFY COLUMN code VARCHAR(255) DEFAULT NULL;
ALTER TABLE colors ADD INDEX code_index(code);

COMMIT;