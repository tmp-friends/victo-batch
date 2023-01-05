#!/bin/sh
gcloud functions deploy \
  insertFanartTweets \
  --region=asia-northeast1 \
  --runtime=nodejs16 \
  --source=. \
  --build-env-vars-file=.env.yaml \
  --entry-point=insertFanartTweets \
  --trigger-http
