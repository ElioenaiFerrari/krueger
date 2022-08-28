FROM golang:1.18-bullseye
RUN mkdir /app
WORKDIR /app

COPY go.* .
RUN go mod tidy
COPY . .
RUN go build -ldflags '-s -w' -o bin/main
VOLUME [ "./krueger.yaml" ]

CMD [ "./bin/main" ]