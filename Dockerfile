FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY commands.go commands.go

RUN go build cmd/main.go

FROM ubuntu:jammy
COPY --from=builder /app/main /connector

ENTRYPOINT [ "/connector" ]
CMD [ "serve" ]
