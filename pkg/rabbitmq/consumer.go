package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type SubscriptionExchange struct {
	ExchangeName    string
	ExchangeType    string
	QueueName       string
	NumberOfWorkers int
}

func (c *AmqpClient) Subscribe(args *SubscriptionExchange, handler func(delivery *amqp.Delivery)) {
	args.Validate()

	if c.Connection == nil {
		panic("Cannot initialize the connection broker. Please initialize the connection first.")
	}

	ch, err := c.Connection.Channel()
	if err != nil {
		panic("Failed to open a channel: " + err.Error())
	}

	d, err := ch.Consume(args.QueueName, "", false, false, false, false, nil)
	if err != nil {
		logrus.Error("Failed to register a consumer: " + err.Error())
	}

	forever := make(chan bool)

	go func() {
		for msg := range d {
			logrus.Infof("Received a message: %s", msg.Body)
			handler(&msg)
		}
	}()

	<-forever
}

func (c SubscriptionExchange) Validate() {
	if c.ExchangeName == "" {
		panic("Exchange name cannot be empty")
	}

	if c.ExchangeType == "" {
		panic("Exchange type cannot be empty")
	}

	if c.QueueName == "" {
		panic("Queue name cannot be empty")
	}

	if c.NumberOfWorkers <= 0 {
		panic("Number of workers must be greater than 0")
	}
}
