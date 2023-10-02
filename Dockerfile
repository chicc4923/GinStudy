FROM golang:1.21.1-alpine

ENV GOPROXY https://goproxy.io
#ENV GO111MODULE ON
ADD . /app
WORKDIR /app

RUN go mod tidy

CMD go run main.go
EXPOSE 8080