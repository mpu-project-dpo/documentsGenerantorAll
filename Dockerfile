FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on \
    GOARCH=amd64 \

RUN apk add --no-cache git

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o main .

FROM alpine

COPY --from=builder /build/main /

COPY config /config

USER nobody

CMD ["/main"]