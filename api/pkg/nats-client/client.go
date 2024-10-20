package natsClient

import "github.com/nats-io/nats.go"

func (c *Config) New() *Client {
	nc, err := nats.Connect(c.URL)
	if err != nil {
		panic(err)
	}

	defer func(nc *nats.Conn) {
		err = nc.Drain()
		if err != nil {
			panic(err)
		}
	}(nc)

	return &Client{nc, c.Topic}
}

func (c *Client) Publish(data []byte) error {
	return c.conn.Publish(c.t, data)
}
