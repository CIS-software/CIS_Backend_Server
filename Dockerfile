FROM golang:1.18-alpine3.15

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main ./cmd/main/

EXPOSE 8080

CMD ["/app/main"]
