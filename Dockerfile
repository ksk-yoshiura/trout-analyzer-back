FROM golang:1.16.3-buster
# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app
# 必要なパッケージをインストール
RUN go mod tidy
# Airをインストール TODO:devではいらない
RUN go install github.com/cosmtrek/air@v1.27.3
# migration
RUN go get github.com/golang-migrate/migrate
# ポート設定
EXPOSE 3000
# 実行
CMD go run main.go router.go