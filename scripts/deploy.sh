#!/bin/sh
# Tips: 開発環境のPlanetScaleはローカルでconnectしておかないと"Can't reach server"エラーになる
rm -rf dist

# docker-compose up でnode_modulesの名前付きボリュームを作っておく
docker exec -it \
  victo-batch \
  sh -c "npm run build"

gcloud functions deploy \
  insertFanartTweets \
  --region=asia-northeast1 \
  --runtime=nodejs16 \
  --source=. \
  --entry-point=insertFanartTweets \
  --trigger-http
