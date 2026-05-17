package testmain

import (
	"testing"
	"trout-analyzer-back/models"

	"github.com/stretchr/testify/assert"
)

/**
  フィールド一覧取得メソッドテスト
	仮テスト
*/
func TestGetAllUsers(t *testing.T) {
	if testing.Short() {
		// スキップ時のメッセージ
		t.Skip("skipping this test")
	}
}

/**
  フィールド一バリデーションテスト
*/
func TestUserValidate(t *testing.T) {
	if testing.Short() {
		// スキップ時のメッセージ
		t.Skip("skipping this test")
	}
}

func TestFindUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping this test")
	}
	t.Run("success FindUser()", func(t *testing.T) {
		u := models.User{Email: "nakata@example.com"}
		user := models.FindUser(u)

		assert.NotEqual(t, models.User{}, user, "failed FindUser()")
		assert.Equal(t, u.Email, user.Email)

		t.Logf("user: %+v", user)
		t.Logf("user.Password: %s", user.Password)
		t.Logf("user.Email: %s", user.Email)
	})

}

func TestCreateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping this test")
	}

	t.Run("success CreateUser()", func(t *testing.T) {
		u := models.User{}
		u.Password = "password"
		u.Email = "nakata@example.com"
		err := models.CreateUser(u)

		assert.NoError(t, err, "failed CreateUser()")

		created := models.FindUser(models.User{Email: u.Email})
		assert.NotEqual(t, models.User{}, created, "failed CreateUser()")
		assert.Equal(t, u.Email, created.Email)

		t.Logf("user: %+v", created)
		t.Logf("user.Email: %s", created.Email)
	})
}
