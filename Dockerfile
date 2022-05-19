# syntax=docker/dockerfile:1
FROM golang:1.18.1-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ENV USER=go
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/cis_backend_server/

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/cis_backend_server ./cmd/main
COPY /config /go/bin

#EXPOSE 8080
#
#CMD ["/cis_backend_server"]

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/bin/config.toml /go/bin/config.toml
COPY --from=builder /go/bin/cis_backend_server /go/bin/cis_backend_server

USER go:go

EXPOSE 8080

ENTRYPOINT ["/go/bin/cis_backend_server"]
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#
#WORKDIR /root/
#
#COPY --from=0 /go/src/cis_backend_server/app ./
#
##COPY app ./cmd/main
#
#CMD ["./app"]
## build stage
#FROM golang:1.18.1-alpine as builder
#
##RUN apk add --no-cache git
##RUN apk --no-cache add ca-certificates
#
#ENV GO111MODULE=on
#
##RUN addgroup -S go \
##    && adduser -S -u 10000 -g go go
##RUN addgroup -g 2000 go && adduser -u 2000 -G go -s /bin/sh -D go
##RUN chown -R go:go /home/go/
##RUN useradd -ms /bin/bash golang
##RUN addgroup -g 2000 golang \ && adduser -u 2000 -G golang -s /bin/sh -D golang
#
#WORKDIR /src
##COPY . /home/go/cis-backend-server
#
##WORKDIR /home/go/cis-backend-server
#
#COPY go.mod go.sum ./
#RUN go mod download
#
##USER go
#
##COPY /cmd/main /cis-backend-server
##RUN go build -o /cis-backend-server
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app ./cmd/main
##COPY --chown=golang:golang . .
#
### final stage
#FROM alpine
##FROM alpine:latest
#COPY --from=builder /app /app
##COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
##COPY --from=0 /etc/passwd /etc/passwd
##COPY --from=builder /home/go/cis-backend-server /home/go/cis-backend-server
##COPY /app /app
##USER go
#
#EXPOSE 8080
#ENTRYPOINT ["/app"]

## allows app_env to be set during build (defaults to empty string)
#ARG app_env
## sets an environment variable to app_env argument, this way the variable     will persist in the container for use in code
#ENV APP_ENV $app_env
#
#COPY ./ /go/src/github.com/CIS-software/CIS_Backend_Server/
#
#WORKDIR /go/src/github.com/CIS-software/CIS_Backend_Server/cmd/main
#
## install all dependencies
#RUN go get ./...
#
## build the binary
#RUN go build -o /cis_backend_server
#
## final stage
#FROM scratch
#
#COPY --from=build-env /cis_backend_server /
#
## Put back once we have an application
#CMD ["/cis_backend_server"]
#
#EXPOSE 8080

# build stage
#FROM golang:1.18.1-alpine as builder
#
#RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
#
#ENV GO111MODULE=on
#ENV USER=appuser
#ENV UID=10001
#
#RUN adduser \
#    --disabled-password \
#    --gecos "" \
#    --home "/nonexistent" \
#    --shell "/sbin/nologin" \
#    --no-create-home \
#    --uid "${UID}" \
#    "${USER}"
#
##RUN addgroup -g 2000 golang && adduser -u 2000 -G golang -s /bin/sh -D golang
##RUN chown -R golang:golang /home/golang/
#
#WORKDIR $GOPATH/src/mypackage/cis_backend_server/
##COPY . .
#
##RUN go get -d -v
##RUN go get ./...
#
#COPY go.mod go.sum ./
#RUN go mod download
##RUN go mod download
##RUN go mod verify
#
##RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello
##RUN GOOS=linux GOARCH=amd64 go build -v -o /go/bin/cis_backend_server ./cmd/main
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/cis_backend_server ./cmd/main
#
## final stage
#FROM scratch
#
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/group /etc/group
#
#COPY --from=builder /go/bin/cis_backend_server /go/bin/cis_backend_server
#
#USER appuser:appuser
#
#EXPOSE 8080
#ENTRYPOINT ["/go/bin/cis_backend_server"]