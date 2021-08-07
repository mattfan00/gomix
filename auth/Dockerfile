FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY auth .
COPY utl .

RUN go build -o /bin/main cmd/auth/main.go

EXPOSE 8080

CMD [ "/bin/main" ]

