package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

type Image struct {
	Image string `json:"image"`
}

/** セッションを返す */
func createSession() *session.Session {
	// 読み込み
	err := godotenv.Load()
	if err != nil { // 本番環境には.envは置いていない
		godotenv.Load("./backend/.env.prod")
	}
	S3_REGION := os.Getenv("S3_REGION")
	S3_ENDPOINT := os.Getenv("S3_ENDPOINT")
	// 特に設定しなくても環境変数にセットしたクレデンシャル情報を利用して接続してくれる
	cfg := aws.Config{
		Region:   aws.String(S3_REGION),
		Endpoint: aws.String(S3_ENDPOINT), // コンテナ内からアクセスする場合はホストをサービス名で指定
		// S3ForcePathStyle: aws.Bool(true),          // ローカルで動かす場合は必須
	}
	return session.Must(session.NewSession(&cfg))
}

/** 画像名を現在時刻からハッシュ値で作成 */
func CreateImageName() string {
	// 現在時刻
	current_time := time.Now()
	s := current_time.String()
	b := []byte(s)
	sha256 := sha256.Sum256(b)
	return hex.EncodeToString(sha256[:])
}

/** S3にアップロード */
func UploadToS3(image Image, image_file string) {
	// 読み込み
	err := godotenv.Load()
	if err != nil { // 本番環境には.envは置いていない
		godotenv.Load("./backend/.env.prod")
	}
	S3_BUCKET := os.Getenv("S3_BUCKET")
	fmt.Println(S3_BUCKET)
	// セッション作成
	sess := createSession()

	image_base64 := image.Image
	b64data := image_base64[strings.IndexByte(image_base64, ',')+1:]
	// デコード
	decode, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		log.Fatal(err)
	}

	// ファイル名
	file_names := []string{
		image_file + ".webp",
		image_file + ".png",
	}

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	for _, file_name := range file_names {
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(S3_BUCKET),
			Key:    aws.String(file_name),
			Body:   bytes.NewReader(decode),
		})
		if err != nil {
			fmt.Printf("when upload :%s", err)
			log.Fatal(err)
		}
	}
	log.Println("done")

}
