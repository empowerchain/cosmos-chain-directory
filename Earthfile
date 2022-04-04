VERSION 0.6
FROM golang:1.16-alpine
WORKDIR /app

project-files:
    COPY go.mod go.sum .
    COPY *.go .
    COPY chain-registry ./chain-registry

build:
    FROM +project-files
    RUN go build -o dist/server
    SAVE ARTIFACT dist/server dist/server AS LOCAL dist/server

docker:
    FROM alpine
    WORKDIR /app
    COPY +build/dist/server .
    CMD [ "/app/server" ]
    SAVE IMAGE --push empowergjermund/cosmos-chain-directory:latest