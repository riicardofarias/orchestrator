package stack

import (
	"context"
	"github.com/google/uuid"
	"log"
	"orchestrator/internal/domain/stack"
	"time"
)

type (
	CreateInput struct {
		Name        string
		Description string
	}

	CreateOutput struct {
		ID string `json:"id"`
	}

	CreateUseCase interface {
		Handle(ctx context.Context, input CreateInput) (CreateOutput, error)
	}

	stackUseCase struct {
		stackRepository stack.RepositoryInterface
	}
)

func NewCreateUseCase(stackRepository stack.RepositoryInterface) CreateUseCase {
	return &stackUseCase{
		stackRepository: stackRepository,
	}
}

func (s stackUseCase) Handle(ctx context.Context, input CreateInput) (CreateOutput, error) {
	var (
		output CreateOutput
		err    error
	)

	inputEntity := stack.New(
		uuid.NewString(),
		input.Name,
		input.Description,
		time.Now(),
	)

	if err := inputEntity.IsValid(); err != nil {
		return output, err
	}

	id, err := s.stackRepository.Create(ctx, inputEntity)
	if err != nil {
		return output, err
	}

	log.Println("Stack created with ID:", id)

	return CreateOutput{ID: id}, nil
}
