FROM golang:1.26.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api ./cmd/app/main.go

EXPOSE 8000

CMD ["./api"]