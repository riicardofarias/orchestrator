package main

import "orchestrator/internal/infra/adapters/rabbitmq"

func main() {
	adapter := rabbitmq.New(rabbitmq.WithRetries(3))

	if err := adapter.Connect(); err != nil {
		panic(err)
	}
}
