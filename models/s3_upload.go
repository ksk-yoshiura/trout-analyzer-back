package models

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func newS3() (*s3.S3, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	ak := "minio"
	sk := "minio123"
	cfg := aws.Config{
		Credentials:      credentials.NewStaticCredentials(ak, sk, ""),
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         aws.String("http://127.0.0.1:9090"),
		S3ForcePathStyle: aws.Bool(true),
	}
	return s3.New(s, &cfg), nil
}

// セッションを返す
func createSession() *session.Session {
	// 特に設定しなくても環境変数にセットしたクレデンシャル情報を利用して接続してくれる
	cfg := aws.Config{
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         aws.String("http://minio:9000"), // コンテナ内からアクセスする場合はホストをサービス名で指定
		S3ForcePathStyle: aws.Bool(true),                  // ローカルで動かす場合は必須
	}
	return session.Must(session.NewSession(&cfg))
}

func upload_minio() {
	sess := createSession()

	// ファイルを開く
	targetFilePath := "./sample.txt"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucketName := "develop"
	objectKey := "test3"

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")

}

func upload() {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "test",
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Region:                        aws.String("ap-northeast-1"),
		},
	}))

	// ファイルを開く
	targetFilePath := "./sample.txt"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucketName := "trout-analyzer-upload"
	objectKey := "test3"

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
