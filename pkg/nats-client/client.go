package nats_client

import "github.com/nats-io/nats.go"

func (c *Config) New() *nats.Conn {
	nc, err := nats.Connect(c.Url)
	if err != nil {
		panic(err)
	}

	defer nc.Drain()

	return nc
}
