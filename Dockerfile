FROM golang:alpine3.11

RUN apk add --update --no-cache bash gcc git make musl-dev
