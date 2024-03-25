package stack

import (
	"orchestrator/internal/domain/stack"
	"time"
)

type Stack struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description,omitempty"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}

func NewFromDomain(stack stack.Stack) *Stack {
	return &Stack{
		ID:          stack.ID,
		Name:        stack.Name,
		Description: stack.Description,
		CreatedAt:   stack.CreatedAt,
		UpdatedAt:   stack.UpdatedAt,
	}
}

func (s *Stack) ToDomain() stack.Stack {
	return stack.Stack{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}
