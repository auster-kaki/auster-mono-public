#!/bin/sh

# RabbitMQをバックグラウンドで起動
rabbitmq-server -detached &

# RabbitMQが起動するまで待機

sleep 5  # 必要に応じて時間を調整してください

# Pythonスクリプトを並列で実行

python rapp.py &
python rworker.py &

# フォアグラウンドでプロセスを維持
wait   