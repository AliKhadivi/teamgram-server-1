FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN ./scripts/build.sh

FROM ubuntu:latest
WORKDIR /app
RUN apt update && apt install ffmpeg
COPY --from=builder /app/teamgramd/ /app/
ENTRYPOINT /app/docker/entrypoint.sh




