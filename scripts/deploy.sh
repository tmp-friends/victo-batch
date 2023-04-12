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
# タイムアウトが1hに延びるため第2世代にする
gcloud beta functions deploy \
  $target_func \
  --gen2 \
  --trigger-http \
  --region asia-northeast1 \
  --runtime go119 \
  --memory 128MiB \
  --source ./functions \
  --entry-point $entry_point \
  --set-env-vars SOURCE_DIR=serverless_function_source_code/
