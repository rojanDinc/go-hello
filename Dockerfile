FROM golang:1.17-alpine AS builder

WORKDIR /build

COPY go.mod .
RUN go mod download
COPY *.go .

RUN go build -o app

# ---
FROM alpine:3.15

WORKDIR /

COPY --from=builder /build/app app

ENTRYPOINT ["./app"]
