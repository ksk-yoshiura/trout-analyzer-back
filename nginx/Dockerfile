FROM nginx:1.17.4-alpine
# インクルード用のディレクトリ内を削除
RUN rm -f /etc/nginx/conf.d/*
# Nginxの設定ファイルをコンテナにコピー
ADD ./nginx/*.conf /etc/nginx/conf.d/default.conf
# ポート設定
EXPOSE 80
# ビルド完了後にNginxを起動
CMD /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf
