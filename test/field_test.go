package testmain

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 開始処理
	log.Print("setup")
	// パッケージ内のテストの実行
	code := m.Run()
	// 終了処理
	log.Print("tear-down")
	// テストの終了コードで exit
	os.Exit(code)
}

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
