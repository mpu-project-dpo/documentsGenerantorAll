package consumer

import "github.com/nats-io/nats.go"

type Config struct {
	URL   string
	Topic string
}

type Client struct {
	conn *nats.Conn
	t    string
}
