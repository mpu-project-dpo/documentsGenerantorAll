package producer

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func (c *Config) New() *Client {
	nc, err := nats.Connect(c.URL)
	if err != nil {
		panic(err)
	}
	if nc.Status() != nats.CONNECTED {
		panic(fmt.Sprintf("Connection to Nats is aborted status %d", nc.Status()))
	}

	return &Client{nc, c.Topic}
}

func (c *Client) Publish(data []byte) error {
	return c.conn.Publish(c.t, data)
}
