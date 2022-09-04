package main

import (
	"fmt"
	"log"
	"os"

	// _ "github.com/aws/aws-sdk-go/aws"
	// _ "github.com/aws/aws-sdk-go/aws/session"
	// _ "github.com/aws/aws-sdk-go/service/ssm"
	"github.com/joho/godotenv"
	// _ "github.com/go-sql-driver/mysql"
)

func getDatabasePassword() string {

	// sess, err := session.NewSessionWithOptions(session.Options{
	// 	Config:  aws.Config{Region: aws.String("ap-northeast-1")},
	// 	Profile: "default",
	// })
	// if err != nil {
	// 	fmt.Println(err)

	// 	// log.Fatal(err)
	// }
	// svc := ssm.New(sess)
	// fmt.Println(svc)

	// res, err := svc.GetParameter(&ssm.GetParameterInput{
	// 	Name:           aws.String("/tranaza/DB_PASSWORD"),
	// 	WithDecryption: aws.Bool(true),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	DB_PASSWORD := "test" // *res.Parameter.Value
	return DB_PASSWORD
}

func getDBConfig() string {
	err := godotenv.Load("./backend/.env.prod")

	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	password := getDatabasePassword()

	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"

	return CONNECT
}

func executeInitialize() {
	CONNECT := getDBConfig()

	fmt.Println(CONNECT)

	// db, err := sql.Open("mysql", CONNECT)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Ping()

	// if err != nil {
	// 	fmt.Println("データベース接続失敗")
	// } else {
	// 	fmt.Println("データベース接続成功")
	// }

}

func main() {
	fmt.Println("test in progress")
	executeInitialize()
}
