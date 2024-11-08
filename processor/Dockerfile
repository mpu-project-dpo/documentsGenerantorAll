FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on \
    GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /src

COPY . .

WORKDIR /src/processor

RUN go mod download
RUN go build -o main .

FROM alpine:3.20

COPY --from=builder /src/processor/main /

COPY processor/config /config

USER nobody

HEALTHCHECK --interval=5s --timeout=10s --retries=3 CMD curl -sS 127.0.0.1:8000 || exit 1

CMD ["/main"]