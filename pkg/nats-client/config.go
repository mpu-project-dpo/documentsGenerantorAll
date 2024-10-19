package nats_client

import "github.com/nats-io/nats.go"

type Config struct {
	Url   string
	Topic string
}

type Client struct {
	conn *nats.Conn
	t    string
}
