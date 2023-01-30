package controllers

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"trout-analyzer-back/models"

	jwt "github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	UID   int    `json:"uid"`
	Email string `json:"email"`

	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

/**
 * サインアップ
 */
func Signup(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// バリデーション
	if err := c.Validate(user); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			c.Logger().Error(k + ": " + err.Error())
		}
		return err
	}

	if u := models.FindUser(models.User{Email: user.Email}); u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "email already exists",
		}
	}

	// ユーザーパスワード暗号化
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	models.CreateUser(user)
	// パスワードは空にする
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}

/**
 * ログイン
 */
func Login(c echo.Context) error {
	// 入力レコード
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	// メールアドレスからユーザ-レコード取得
	user := models.FindUser(models.User{Email: u.Email})
	// パスワードチェック
	match := CheckPasswordHash(user.Password, u.Password)

	if user.ID == 0 || !match { // 既存ユーザーおよびパスワードが合致するか
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid email or password",
		}
	}

	claims := &jwtCustomClaims{
		int(user.ID),
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

/**
 * パスワード再設定
 */
func ResetPassword(c echo.Context) error {

	// データセット
	n := models.NewPassword{}
	if err := c.Bind(&n); err != nil {
		return err
	}
	current_password := n.Password

	// ユーザーIDからユーザ-レコード取得
	uid := userIDFromToken(c)
	u := models.User{}
	user := models.GetUser(u, uid)
	// パスワードチェック
	match := CheckPasswordHash(user.Password, current_password)

	if user.ID == 0 || !match { // 現在のパスワードが一致するか
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid password",
		}
	}
	// 新規パスワード
	new_password := n.NewPassword
	// ユーザーパスワード暗号化
	new_hash, _ := HashPassword(new_password)

	// 確認パスワード
	confirm_password := n.ConfirmPassword

	if new_password != confirm_password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "new password should match with the password confirm",
		}
	}

	user.Password = new_hash
	models.UpdateUser(user, uid)

	return c.JSON(http.StatusCreated, u)
}

/**
 * トークンからユーザID取得
 */
func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}

/**
 * パスワードハッシュ化
 */
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/**
 * パスワードチェック
 */
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
