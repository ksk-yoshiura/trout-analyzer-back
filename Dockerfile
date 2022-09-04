FROM golang:1.16.3-buster
# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app
# 必要なパッケージをインストール
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3
# godotenv
RUN go install github.com/joho/godotenv
# SDK
RUN go install github.com/aws/aws-sdk-go

# AWS CLI
ARG GLIBC_VERSION=2.35-r0
ARG AWSCLI_VERSION=2.6.1

# install glibc compatibility for alpine
RUN apk --no-cache add \
  binutils \
  curl \
  && curl -sL https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub -o /etc/apk/keys/sgerrand.rsa.pub \
  && curl -sLO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk \
  && curl -sLO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk \
  && curl -sLO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-i18n-${GLIBC_VERSION}.apk \
  && apk add --no-cache \
  glibc-${GLIBC_VERSION}.apk \
  glibc-bin-${GLIBC_VERSION}.apk \
  glibc-i18n-${GLIBC_VERSION}.apk \
  && /usr/glibc-compat/bin/localedef -i en_US -f UTF-8 en_US.UTF-8 \
  && ln -sf /usr/glibc-compat/lib/ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2 \
  && curl -sL https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWSCLI_VERSION}.zip -o awscliv2.zip \
  && unzip awscliv2.zip \
  && aws/install \
  && rm -rf \
  awscliv2.zip \
  aws \
  /usr/local/aws-cli/v2/current/dist/aws_completer \
  /usr/local/aws-cli/v2/current/dist/awscli/data/ac.index \
  /usr/local/aws-cli/v2/current/dist/awscli/examples \
  glibc-*.apk \
  && find /usr/local/aws-cli/v2/current/dist/awscli/botocore/data -name examples-1.json -delete \
  && apk --no-cache del \
  binutils \
  curl \
  && rm -rf /var/cache/apk/*

# ポート設定
EXPOSE 3000
# 実行
CMD go run main.go router.go