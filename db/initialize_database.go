package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func getDBConfig() string {
	// 読み込み
	err := godotenv.Load("./backend/.env.prod")
	if err != nil {
		log.Fatal(err)
	}
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	password := os.Getenv("DB_PASSWORD")

	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"

	return CONNECT
}

func executeInitialize() {
	CONNECT := getDBConfig()
	db, err := sql.Open("mysql", CONNECT)
	defer db.Close()

	fmt.Println(CONNECT)
	// out, err := exec.Command("initialize_db_sql.sh").Output()
	// out, err := exec.Command("migrate", "-version").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")
	}
}

func main() {
	executeInitialize()
}
