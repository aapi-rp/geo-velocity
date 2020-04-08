FROM golang:1.12-alpine


RUN apk update \
    && apk add sqlite
