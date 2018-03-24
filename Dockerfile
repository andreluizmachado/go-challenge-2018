FROM golang:1.10-alpine

RUN  apk update \
    && apk add build-base
