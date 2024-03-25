package job

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"orchestrator/internal/domain/job"
	"orchestrator/internal/domain/stack"
	"orchestrator/internal/domain/step"
	"orchestrator/pkg/rabbitmq"
)

var (
	ErrStepNotFound        = errors.New("step_not_found")
	ErrStepAlreadyFinished = errors.New("step_already_finished")
)

type (
	HandleJobResponse struct {
		UUID    string `json:"uuid"`
		Payload string `json:"payload"`
		Success bool   `json:"success"`
	}

	HandleJobUseCase interface {
		Handle(ctx context.Context) error
	}

	handleJobUseCase struct {
		jobRepository   job.RepositoryInterface
		stackRepository stack.RepositoryInterface
		stepRepository  step.RepositoryInterface
		amqpClient      *rabbitmq.AmqpClient
	}
)

func NewHandleJobUseCase(
	jobRepository job.RepositoryInterface,
	stackRepository stack.RepositoryInterface,
	stepRepository step.RepositoryInterface,
	amqpClient *rabbitmq.AmqpClient,
) HandleJobUseCase {
	return &handleJobUseCase{
		jobRepository:   jobRepository,
		stackRepository: stackRepository,
		stepRepository:  stepRepository,
		amqpClient:      amqpClient,
	}
}

func (u *handleJobUseCase) Handle(ctx context.Context) error {
	args := &rabbitmq.SubscriptionExchange{
		ExchangeName:    "saga.topic",
		ExchangeType:    "direct",
		QueueName:       "saga.handler",
		NumberOfWorkers: 10,
	}

	u.amqpClient.Subscribe(args, func(delivery *amqp.Delivery) {
		var response HandleJobResponse

		if err := json.Unmarshal(delivery.Body, &response); err != nil {
			logrus.Errorf("failed to unmarshal request: %v", err)
			_ = delivery.Nack(false, false)
		}

		if err := u.handleJobResponse(ctx, response); err != nil {
			logrus.Errorf("failed to handle job response: %v", err)
			_ = delivery.Nack(false, false)
		}

		_ = delivery.Ack(false)
	})

	return nil
}

func (u *handleJobUseCase) handleJobResponse(ctx context.Context, response HandleJobResponse) error {
	stepEntity, err := u.stepRepository.FindByID(ctx, response.UUID)
	if err != nil {
		return err
	}

	if !stepEntity.Exists() {
		return ErrStepNotFound
	}

	if stepEntity.IsFinished() {
		return ErrStepAlreadyFinished
	}

	if response.Success {
		return u.stepRepository.UpdateStatus(ctx, stepEntity.ID, step.StatusFinished)
	} else {
		return u.stepRepository.UpdateStatus(ctx, stepEntity.ID, step.StatusFailed)
	}
}
