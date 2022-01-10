#build stage
FROM golang:rc-alpine AS builder
RUN apk add build-base