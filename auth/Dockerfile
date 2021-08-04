FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/auth cmd/auth
COPY svc/auth svc/auth
COPY pkg pkg

RUN go build -o /bin/main cmd/auth/main.go

EXPOSE 8080

CMD [ "/bin/main" ]

