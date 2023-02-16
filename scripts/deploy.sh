#!/bin/bash
set -eu

help_message() {
  cat << EOF
[対象関数]
  - insertFanartTweets
  - registerVtubers
  - registerHashtags
EOF
}

if [ $# -ne 1 ]; then
  echo "パラメータは1つ必要です"
  echo "./scripts/deploy.sh [対象関数]"
  echo ""
  help_message
  exit 1
fi;

target_func=$1
entry_point=${target_func^}

# Tips: 開発環境のPlanetScaleはローカルでconnectしておかないと"Can't reach server"エラーになる
gcloud functions deploy \
  $target_func \
  --region=asia-northeast1 \
  --runtime=go119 \
  --set-env-vars SOURCE_DIR=serverless_function_source_code/ \
  --source=./functions \
  --entry-point=$entry_point \
  --trigger-http
