package job

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"orchestrator/internal/domain/job"
	"orchestrator/internal/domain/stack"
	"orchestrator/internal/domain/step"
	"orchestrator/pkg/rabbitmq"
)

var (
	ErrStackNotFound = errors.New("stack_not_found")
)

type (
	StartJobRequest struct {
		StackID string
		Payload string
		Queue   string
	}

	StartJobUseCase interface {
		Handle(ctx context.Context, input StartJobRequest) error
	}

	startJobUseCase struct {
		jobRepository   job.RepositoryInterface
		stackRepository stack.RepositoryInterface
		stepRepository  step.RepositoryInterface
		amqpClient      *rabbitmq.AmqpClient
	}
)

func NewStartJobUseCase(
	jobRepository job.RepositoryInterface,
	stackRepository stack.RepositoryInterface,
	stepRepository step.RepositoryInterface,
	amqpClient *rabbitmq.AmqpClient,
) StartJobUseCase {
	return &startJobUseCase{
		jobRepository:   jobRepository,
		stackRepository: stackRepository,
		stepRepository:  stepRepository,
		amqpClient:      amqpClient,
	}
}

func (s *startJobUseCase) Handle(ctx context.Context, input StartJobRequest) error {
	stackEntity, err := s.stackRepository.FindByID(ctx, input.StackID)
	if err != nil {
		return err
	}

	if !stackEntity.Exists() {
		return ErrStackNotFound
	}

	jobEntity := job.NewJob(stackEntity.ID)
	if _, err := s.jobRepository.Create(ctx, jobEntity); err != nil {
		return err
	}

	stepEntity := step.NewStep(
		input.Payload,
		input.Queue,
		stackEntity.ID,
		jobEntity.ID,
	)

	if _, err := s.stepRepository.Create(ctx, stepEntity); err != nil {
		return err
	}

	payload, _ := json.Marshal(input.Payload)
	topic := "saga.topic"

	if err := s.amqpClient.Publish(topic, input.Queue, payload); err != nil {
		log.Println("Failed to publish message: ", err.Error())
		return err
	}

	return nil
}
