#!/bin/sh
gcloud functions deploy \
  insertFanartTweets \
  --region=asia-northeast1 \
  --runtime=nodejs16 \
  --source=. \
  --entry-point=insertFanartTweets \
  --trigger-http
