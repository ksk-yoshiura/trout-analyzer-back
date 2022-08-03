FROM golang:1.16.3-buster

# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app

# 必要なパッケージをインストール
RUN go mod tidy

# Check and Build
RUN go get golang.org/x/lint/golint && \
  make validate && \
  make build-linux

### If use TLS connection in container, add ca-certificates following command.
### > RUN apk add --no-cache ca-certificates
FROM gcr.io/distroless/base-debian10
COPY --from=build-env /app/main /
EXPOSE 80
ENTRYPOINT ["/main"]

# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
CMD ["air"]