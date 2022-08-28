package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetDB returns database connection
func GetDBConn() *gorm.DB {
	CONNECT := GetDBConfig()
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDBConfig() string {
	// 読み込み
	err := godotenv.Load("./backend/.env.prod")
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"
	return CONNECT
}
