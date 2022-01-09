package models

import (
	"trout-analyzer-back/database"

	"time"
)

// Model
type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"first_name"`
	LasttName string `json:"last_name"`
	DeleteFlg string `json:"delete_flg"`
	Group_id  int    `json:"group_id"`
}

var users []User

/** ユーザー一覧取得 */
func (u *User) GetUsers() (err error) {
	db := database.GetDBConn()
	return db.Find(&users).Error
}

// /** ユーザー作成 */
// func (u *User) Create() (err error) {
// 	db := database.GetDB()
// 	return db.Create(u).Error
// }

// /** ユーザー取得 */
// func (u *User) FindByID(id uint) (err error) {
// 	db := database.GetDB()
// 	return db.Where("id = ?", id).First(u).Error
// }
