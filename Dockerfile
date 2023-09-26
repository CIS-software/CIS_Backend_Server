# syntax=docker/dockerfile:1
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/cis_backend_server/
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/cis_backend_server ./cmd

FROM alpine
COPY --from=builder /go/bin/cis_backend_server /cis_backend_server/cis_backend_server
COPY /config /cis_backend_server/config

EXPOSE 8080
ENTRYPOINT ["/cis_backend_server/cis_backend_server"]