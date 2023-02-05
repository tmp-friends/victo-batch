#!/bin/sh
# Tips: 開発環境のPlanetScaleはローカルでconnectしておかないと"Can't reach server"エラーになる
gcloud functions deploy \
  registerVtubers \
  --region=asia-northeast1 \
  --runtime=go119 \
  --set-env-vars SOURCE_DIR=serverless_function_source_code/ \
  --source=./functions \
  --entry-point=RegisterVtubers \
  --trigger-http
