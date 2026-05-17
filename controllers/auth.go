package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"trout-analyzer-back/models"

	validation "github.com/go-ozzo/ozzo-validation"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	UID   int    `json:"uid"`
	Email string `json:"email"`

	jwt.RegisteredClaims
}

var signingKey = getSigningKey()

func getSigningKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return []byte(key)
}

// JWTMiddleware はリクエストのAuthorizationヘッダーからJWTトークンを検証するミドルウェア
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "missing authorization header"}
			}
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid authorization header format"}
			}
			claims := &jwtCustomClaims{}
			token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "unexpected signing method"}
				}
				return signingKey, nil
			})
			if err != nil || !token.Valid {
				return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid or expired token"}
			}
			c.Set("user", token)
			return next(c)
		}
	}
}

/**
 * サインアップ
 */
func Signup(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// メールアドレスからユーザ-レコード取得
	user := models.FindUser(models.User{Email: u.Email})
	fmt.Printf("Email2: %s", user.Email)
	// パスワードチェック
	match := CheckPasswordHash(user.Password, u.Password)

	if user.ID == 0 || !match { // 既存ユーザーおよびパスワードが合致するか
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid email or password",
		}
	}

	claims := &jwtCustomClaims{
		UID:   int(user.ID),
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
