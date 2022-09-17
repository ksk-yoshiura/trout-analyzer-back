FROM golang:1.16.3-buster
# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app
# 必要なパッケージをインストール
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz

# # godotenv
# RUN go install github.com/joho/godotenv
# # SDK
# RUN go install github.com/aws/aws-sdk-go
# # mysqlドライバ
# RUN go install github.com/go-sql-driver/mysql
# # migration
# RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ### EXEC ECSでAWSコマンドを使うのに必須
# # 前提パッケージのインストール
# RUN apt-get update && apt-get install -y less vim curl unzip sudo
# # aws cli v2 のインストール
# # https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/install-cliv2-linux.html
# RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
# RUN unzip awscliv2.zip
# RUN sudo ./aws/install

# ポート設定
EXPOSE 3000
# 実行
CMD go run main.go router.go