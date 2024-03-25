package stack

import (
	"context"
)

type RepositoryInterface interface {
	FindByID(ctx context.Context, id string) (Stack, error)
	Create(ctx context.Context, stack Stack) (string, error)
}
