package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// GetDB returns database connection
func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	DBMS := "mysql"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"
	return DBMS, CONNECT
}
