# syntax=docker/dockerfile:1
FROM golang:1.18.1-alpine
WORKDIR $GOPATH/src/cis_backend_server/
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/cis_backend_server ./cmd

FROM alpine
COPY --from=0 /go/bin/cis_backend_server /go/bin/cis_backend_server
EXPOSE 8080
ENTRYPOINT ["/go/bin/cis_backend_server"]