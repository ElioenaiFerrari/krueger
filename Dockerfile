FROM golang:1.18-alpine

RUN mkdir /app
WORKDIR /app

COPY go.* .
RUN go mod tidy
COPY . .
RUN go build -ldflags '-s -w' -o bin/main

CMD [ "./bin/main" ]