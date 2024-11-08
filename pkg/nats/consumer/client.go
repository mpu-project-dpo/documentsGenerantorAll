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

	_, err = nc.Subscribe(config.Topic, convertToSubFunc[T](fn))
	if err != nil {
		return nil
	}

	return Client{nc}.Run
}

func (c Client) Run() error {
	select {
	//TODO Add graceful shutdown
	}

	err := c.conn.Drain()
	if err != nil {
		return err
	}

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
