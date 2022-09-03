package main

import (
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
)

// func getDatabasePassword() string {
// 	sess, err := session.NewSessionWithOptions(session.Options{
// 		Config:  aws.Config{Region: aws.String("ap-northeast-1")},
// 		Profile: "default",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	svc := ssm.New(sess)

// 	res, err := svc.GetParameter(&ssm.GetParameterInput{
// 		Name:           aws.String("/tranaza/DB_PASSWORD"),
// 		WithDecryption: aws.Bool(true),
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	DB_PASSWORD := *res.Parameter.Value

// 	return DB_PASSWORD
// }

// func getDBConfig() string {
// 	// 読み込み
// 	err := godotenv.Load("./backend/.env.prod")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	user := os.Getenv("DB_USER")
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	database_name := os.Getenv("DB_DATABASE_NAME")

// 	password := getDatabasePassword()
// 	fmt.Println(password)

// 	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"

// 	return CONNECT
// }

func executeInitialize() {
	CONNECT := "testtest"
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
	executeInitialize()
}
