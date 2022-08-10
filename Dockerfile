FROM golang:1.16.3-buster
# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app
# 必要なパッケージをインストール
RUN go mod tidy
# Airをインストール TODO:devではいらない
RUN go install github.com/cosmtrek/air@v1.27.3
# ポート設定
EXPOSE 3000
# TODO:devではair
CMD go run main.go router.go