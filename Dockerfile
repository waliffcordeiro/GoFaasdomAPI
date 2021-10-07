FROM golang:alpine as builder

ENV GO111MODULE=on
RUN apk update && apk add --no-cache git

RUN mkdir app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

CMD ["./bin/main"]