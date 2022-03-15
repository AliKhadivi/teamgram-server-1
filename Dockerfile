FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN ./scripts/build.sh

FROM ubuntu:latest
WORKDIR /app
RUN apt update -y && apt install -y ffmpeg
COPY --from=builder /app/teamgramd/ /app/
ENTRYPOINT /app/docker/entrypoint.sh




