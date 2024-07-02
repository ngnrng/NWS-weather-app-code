FROM golang:1.18-alpine as builder

ARG OW_KEY

ENV OW_KEY=${OW_KEY}

WORKDIR /app

COPY . /app

COPY go.mod  ./

RUN apk update && apk add git

RUN apk --no-cache add ca-certificates

RUN go build -o weather-app .

# Command to run the executable
CMD ("./weather-app")