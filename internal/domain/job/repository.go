package job

import (
	"context"
)

type RepositoryInterface interface {
	Create(ctx context.Context, job Job) (string, error)
	UpdateStatus(ctx context.Context, id, status string) error
}
