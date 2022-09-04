package database

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDatabasePassword() string {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("ap-northeast-1")},
		Profile: "default",
	})
	if err != nil {
		log.Fatal(err)
	}
	svc := ssm.New(sess)

	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/tranaza/DB_PASSWORD"),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}

	DB_PASSWORD := *res.Parameter.Value

	return DB_PASSWORD
}

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
	password := getDatabasePassword()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	CONNECT := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"
	return CONNECT
}
