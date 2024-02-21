FROM golang:1.22-alpine as build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN apk add --no-cache --virtual .build-deps build-base

WORKDIR ${GOPATH}/src/app

COPY . .

RUN go env -w GOPROXY=direct

RUN go build -ldflags '-extldflags "-static"' -o bin/app main.go

CMD ["./bin/app"]