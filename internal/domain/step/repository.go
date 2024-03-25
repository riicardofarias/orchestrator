package step

import (
	"context"
)

type RepositoryInterface interface {
	FindByID(ctx context.Context, id string) (Step, error)
	Create(ctx context.Context, step Step) (string, error)
	UpdateStatus(ctx context.Context, id, status string) error
}
