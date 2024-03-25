package rabbitmq

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type IMessagingClient interface {
	ConnectToBroker(config *Config)
	Publish(exchangeName, exchangeType, routingKey string, message []byte)
}

type AmqpClient struct {
	Connection *amqp.Connection
}

type ExchangeMessage struct {
	Body         []byte
	Exchange     string
	ExchangeType string
}

func New() *AmqpClient {
	return &AmqpClient{}
}

func (c *AmqpClient) ConnectToBroker(config *Config) error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(url)

	if err != nil {
		return err
	}

	logrus.Info("Connected to broker RabbitMQ")

	c.Connection = conn

	return nil
}
