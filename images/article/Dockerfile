FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/article cmd/article
COPY svc/article svc/article
COPY pkg pkg

RUN go build -o /bin/main cmd/article/main.go

EXPOSE 8081

CMD [ "/bin/main" ]
