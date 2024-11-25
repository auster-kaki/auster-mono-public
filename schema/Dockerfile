# ベースイメージとして MySQL を使用
FROM mysql:8.0.40

# タイムゾーン設定
ENV TZ=Asia/Tokyo

# 環境変数設定
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes \
    MYSQL_DATABASE=auster

# MySQL 設定ファイルと初期スクリプトをコピー
COPY ./mysql/my.conf /etc/mysql/my.cnf
COPY ./schema.sql /docker-entrypoint-initdb.d/init.sql

# ポートを開放
EXPOSE 3306

# 健康チェック (MySQL が起動するまで待機)
HEALTHCHECK --interval=10s --timeout=5s --retries=3 CMD mysqladmin ping -h localhost || exit 1
