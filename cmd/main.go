package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"orchestrator/config"
	jobRepo "orchestrator/internal/infra/database/mongodb/job"
	stackRepo "orchestrator/internal/infra/database/mongodb/stack"
	stepRepo "orchestrator/internal/infra/database/mongodb/step"
	jobUC "orchestrator/internal/usecases/job"
	mongo "orchestrator/pkg/mongodb"
	"orchestrator/pkg/rabbitmq"
)

func main() {
	ctx := context.Background()

	config.InitConfigs()

	dbConfig := mongo.GetDatabaseConfig()
	db := mongo.NewMongoDB(mongo.ServerConfig{
		Database: dbConfig.Database,
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
	}, mongo.NewLocalClient())

	client, err := db.GetClient()
	if err != nil {
		panic(err)
	}

	amqpConfig := rabbitmq.GetRabbitMQConfig()
	amqpClient := rabbitmq.New()

	if err := amqpClient.ConnectToBroker(amqpConfig); err != nil {
		panic(err)
	}

	stackRepository := stackRepo.NewRepository(client.Database())
	jobRepository := jobRepo.NewRepository(client.Database())
	stepRepository := stepRepo.NewRepository(client.Database())

	/*
		payload, _ := json.Marshal(map[string]interface{}{
			"name": "Ricardo",
		})

		startJobUC := jobUC.NewStartJobUseCase(jobRepository, stackRepository, stepRepository, amqpClient)
		err = startJobUC.Handle(ctx, jobUC.StartJobRequest{
			StackID: "ddd4d609-7a2f-45a3-af00-0d21e4b61f59",
			Queue:   "order-routing",
			Payload: string(payload),
		})

		if err != nil {
			panic(err)
		}*/

	replyJobUC := jobUC.NewHandleJobUseCase(jobRepository, stackRepository, stepRepository, amqpClient)
	if err := replyJobUC.Handle(ctx); err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(recover.New())

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
