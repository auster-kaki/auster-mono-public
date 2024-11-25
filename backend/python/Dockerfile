# Start of Selection
# syntax=docker/dockerfile:1

# Comments are provided throughout this file to help you get started.
# If you need more help, visit the Dockerfile reference guide at
# https://docs.docker.com/go/dockerfile-reference/

# Want to help us make this template better? Share your feedback here: https://forms.gle/ybq9Krt8jtBL3iCk7

ARG PYTHON_VERSION=3.10.15
FROM python:${PYTHON_VERSION}-slim-bullseye as base

# Pythonがpycファイルを書き込むのを防ぎます。
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

WORKDIR /home/roop

# Create a non-privileged user that the app will run under.
# See https://docs.docker.com/go/dockerfile-user-best-practices/
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /root/.cache/pip to speed up subsequent builds.
# Leverage a bind mount to requirements.txt to avoid having to copy them into
# into this layer.
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y gcc patch g++ wget python3-tk


RUN python -m pip install --upgrade pip
RUN python -m ensurepip --upgrade
RUN python -m pip install --upgrade setuptools
RUN pip install --upgrade setuptools

COPY ./requirements.txt ./

RUN python -m pip install -r requirements.txt


COPY ./roop ./
COPY ./roopapi ./api

RUN wget https://huggingface.co/ezioruan/inswapper_128.onnx/resolve/main/inswapper_128.onnx -O inswapper_128.onnx \
    && mkdir models \
    && mv inswapper_128.onnx ./models/


WORKDIR /home/roop/api
# 書き込み権限を与える
RUN chown -R appuser:appuser /home/roop/api/

# RabbitMQのインストール
COPY mq.sh ./
RUN chmod +x mq.sh && ./mq.sh

RUN apt-get install -y erlang-base \
    erlang-asn1 erlang-crypto erlang-eldap erlang-ftp erlang-inets \
    erlang-mnesia erlang-os-mon erlang-parsetools erlang-public-key \
    erlang-runtime-tools erlang-snmp erlang-ssl \
    erlang-syntax-tools erlang-tftp erlang-tools erlang-xmerl \
    libgl1-mesa-glx ffmpeg rabbitmq-server -y --fix-missing

RUN chown -R appuser:appuser /var/lib/rabbitmq
RUN chown -R appuser:appuser /var/log/rabbitmq
ENV HOME=/var/lib/rabbitmq

COPY ./cmd.sh ./cmd.sh
# スクリプトに実行権限を付与
RUN chmod +x ./cmd.sh

USER appuser
EXPOSE 5050
# Run the application.

# デフォルトのコマンドをスタートアップスクリプトに変更
CMD ["sh", "./cmd.sh"]
