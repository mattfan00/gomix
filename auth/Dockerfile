FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY auth auth
COPY utl utl

RUN go build -o auth/bin/main auth/cmd/main.go

EXPOSE 8080

WORKDIR /app/auth

CMD [ "bin/main" ]

