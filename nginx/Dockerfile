FROM nginx:1.17.4-alpine
# インクルード用のディレクトリ内を削除
RUN rm -f /etc/nginx/conf.d/*
# Nginxの設定ファイルをコンテナにコピー
ADD ./nginx/*.conf /etc/nginx/conf.d/default.conf
# アプリケーションのコピー
COPY ./app /work/app
# ビルド完了後にNginxを起動
CMD /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf