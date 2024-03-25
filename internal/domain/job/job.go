package job

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

const (
	StatusRunning  = "RUNNING"
	StatusFinished = "FINISHED"
	StatusFailed   = "FAILED"
)

type Job struct {
	ID         string    `validate:"required"`
	StartedAt  time.Time `validate:"required"`
	FinishedAt time.Time
	Status     string `validate:"required"`
	StackID    string `validate:"required"`
}

func (j *Job) IsValid() error {
	return validator.New().Struct(j)
}

func NewJob(stackID string) Job {
	return Job{
		ID:        uuid.NewString(),
		Status:    StatusRunning,
		StartedAt: time.Now(),
		StackID:   stackID,
	}
}
