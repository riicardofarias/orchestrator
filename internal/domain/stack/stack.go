package stack

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Stack struct {
	ID          string    `validate:"required"`
	Name        string    `validate:"required"`
	Description string    `validate:"required"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time
}

func New(id, name, description string, createdAt time.Time) Stack {
	return Stack{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func (s *Stack) Exists() bool {
	return s.ID != ""
}

func (s *Stack) IsValid() error {
	return validator.New().Struct(s)
}
