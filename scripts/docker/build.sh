#!/bin/sh
set -eu

# Default param
IMAGE_NAME="tmp-friends/victo-batch"
DOCKERFILE="scripts/docker/Dockerfile"
uid=1000
gid=1000
user=node

cd $(dirname $0)/../..
docker image build -t $IMAGE_NAME:latest -f $DOCKERFILE . \
  --build-arg uid=$uid \
  --build-arg gid=$gid \
  --build-arg user=$user
