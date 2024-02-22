FROM golang:1.22-alpine as build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV GOSUMDB=off

WORKDIR /go/src/app

COPY . .

RUN apk add --no-cache --virtual .build-deps git build-base
RUN go build -ldflags '-extldflags "-static"' -o bin/app main.go

CMD ["./bin/app"]