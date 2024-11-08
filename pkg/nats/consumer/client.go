package consumer

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
)

func New[T any](config *Config, fn func(*T) error) func() error {
	nc, err := nats.Connect(config.URL)
	if err != nil {
		panic(err)
	}

	defer func(nc *nats.Conn) {
		err = nc.Drain()
		if err != nil {
			panic(err)
		}
	}(nc)

	_, err = nc.Subscribe(config.Topic, convertToSubFunc[T](fn))
	if err != nil {
		return nil
	}

	return Client{nc, config.Topic}.Run
}

func (c *Client) Run() error {
	return nil
}

func convertToSubFunc[T any](fn func(*T) error) func(m *nats.Msg) {
	return func(m *nats.Msg) {
		t := new(T)
		err := json.Unmarshal(m.Data, t)
		if err != nil {
			return
		}
		err = fn(t)
		if err != nil {
			return
		}
	}
}
