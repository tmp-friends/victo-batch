FROM node:16-alpine

WORKDIR /opt/batch

COPY package*.json ./

RUN npm install

