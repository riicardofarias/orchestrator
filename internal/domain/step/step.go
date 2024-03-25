package step

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

var (
	StatusRunning  = "RUNNING"
	StatusFailed   = "FAILED"
	StatusFinished = "FINISHED"
)

type Step struct {
	ID         string    `validate:"required"`
	Payload    string    `validate:"required,json"`
	Queue      string    `validate:"required"`
	Status     string    `validate:"required"`
	StartedAt  time.Time `validate:"required"`
	FinishedAt time.Time
	StackID    string `validate:"required"`
	JobID      string `validate:"required"`
}

func (j *Step) IsValid() error {
	return validator.New().Struct(j)
}

func (j *Step) IsRunning() bool {
	return j.Status == StatusRunning
}

func (j *Step) IsFailed() bool {
	return j.Status == StatusFailed
}

func (j *Step) IsFinished() bool {
	return j.Status == StatusFinished
}

func (j *Step) Exists() bool {
	return j.ID != ""
}

func NewStep(payload, queue, stackID, jobID string) Step {
	return Step{
		ID:        uuid.NewString(),
		Payload:   payload,
		Status:    StatusRunning,
		StartedAt: time.Now(),
		Queue:     queue,
		StackID:   stackID,
		JobID:     jobID,
	}
}
