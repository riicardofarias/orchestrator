package main

import (
	"encoding/json"
	"fmt"
	"orchestrator/config"
	"orchestrator/pkg/rabbitmq"
	"time"
)

func main() {

	config.InitConfigs()

	amqpConfig := rabbitmq.GetRabbitMQConfig()
	amqpClient := rabbitmq.New()

	if err := amqpClient.ConnectToBroker(amqpConfig); err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		input, _ := json.Marshal(map[string]interface{}{
			"uuid":   fmt.Sprintf("uuid-%d", i),
			"status": "success",
		})

		if err := amqpClient.Publish("saga.topic", "", input); err != nil {
			panic(err)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
