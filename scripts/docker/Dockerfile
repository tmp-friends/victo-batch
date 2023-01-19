FROM node:16-alpine

USER root
ENV TZ=Asia/Tokyo \
  LANG=ja_JP.UTF-8

ARG uid
ARG gid
ARG user

RUN chown -R $user:$user /var \
  && chmod -R 0755 /var \
  && install -m 0755 -o $user -g $user -d /var/batch \
  && install -m 0755 -o $user -g $user -d /var/batch/node_modules \
  && install -m 0755 -o $user -g $user -d /var/batch/src \
  && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /var/batch

COPY package*.json ./
RUN npm install

USER $user
