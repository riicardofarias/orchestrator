package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (c *AmqpClient) Publish(topicName, routingKey string, message []byte) error {
	if c.Connection == nil {
		panic("Cannot initialize the connection broker. Please initialize the connection first.")
	}

	ch, err := c.Connection.Channel()
	defer func(ch *amqp.Channel) {
		_ = ch.Close()
	}(ch)

	if err != nil {
		panic("Failed to open a channel: " + err.Error())
	}

	err = ch.Publish(
		topicName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Body: message,
		},
	)

	logrus.Infof("A message was sent: %v", string(message))

	return err
}
