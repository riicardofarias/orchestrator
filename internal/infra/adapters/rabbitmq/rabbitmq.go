package rabbitmq

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Adapter interface {
	Connect() error
}

type adapter struct {
	connection *amqp.Connection
}

func New(options ...Option) Adapter {
	opts := &option{}
	for _, option := range options {
		option.apply(opts)
	}

	logrus.Infof("opts: %v", opts.retries)
	return &adapter{}
}

func (c *adapter) Connect() error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", "guest", "guest", "localhost", 5673)

	var err error
	if c.connection, err = amqp.Dial(url); err != nil {
		return err
	}

	logrus.Info("Connected to broker RabbitMQ")
	return nil
}
