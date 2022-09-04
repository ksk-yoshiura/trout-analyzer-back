FROM golang:1.16.3-buster
# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app
# 必要なパッケージをインストール
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3
# godotenv
RUN go install github.com/joho/godotenv@v1.4.0
# SDK
RUN go install github.com/aws/aws-sdk-go@v1.44.91
# ポート設定
EXPOSE 3000
# 実行
CMD go run main.go router.go