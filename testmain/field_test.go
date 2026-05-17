package testmain

import (
	"testing"
)

/**
  フィールド一覧取得メソッドテスト
	仮テスト
*/
func TestGetAllFields(t *testing.T) {
	if testing.Short() {
		// スキップ時のメッセージ
		t.Skip("skipping this test")
	}
}

/**
  フィールド一バリデーションテスト
*/
func TestFieldValidate(t *testing.T) {
	if testing.Short() {
		// スキップ時のメッセージ
		t.Skip("skipping this test")
	}
}
