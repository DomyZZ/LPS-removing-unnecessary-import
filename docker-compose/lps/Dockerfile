FROM --platform=linux/amd64 golang:1.16-alpine

RUN apk add git
RUN apk add gcc
RUN apk add musl-dev
RUN apk add curl
RUN apk add jq

WORKDIR /code

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY connectors ./connectors
COPY http ./http
COPY storage ./storage
COPY pegout ./pegout

RUN go build -o /usr/local/bin/liquidity-provider-server

ARG UID=1000
ARG HOME="/home/lps"
RUN adduser -u "$UID" --home="$HOME" lps -D lps

COPY --chown=lps "docker-compose/lps/config.json" "$HOME/"
COPY --chown=lps "docker-compose/lps/start.sh" "$HOME/"

RUN mkdir -p "$HOME/db"; chown "$UID" "$HOME/db"
RUN mkdir -p "$HOME/logs"; chown "$UID" "$HOME/logs"

WORKDIR "$HOME"

USER lps

EXPOSE 8080
