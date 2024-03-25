package step

import (
	"orchestrator/internal/domain/step"
	"time"
)

type Step struct {
	ID         string    `bson:"_id,omitempty"`
	Payload    string    `bson:"payload"`
	Queue      string    `bson:"queue"`
	StartedAt  time.Time `bson:"started_at"`
	FinishedAt time.Time `bson:"finished_at,omitempty"`
	Status     string    `bson:"status"`
	StackID    string    `bson:"stack_id"`
	JobID      string    `bson:"job_id"`
}

func NewFromDomain(step step.Step) *Step {
	return &Step{
		ID:         step.ID,
		Payload:    step.Payload,
		Queue:      step.Queue,
		StartedAt:  step.StartedAt,
		FinishedAt: step.FinishedAt,
		Status:     step.Status,
		StackID:    step.StackID,
		JobID:      step.JobID,
	}
}

func (s *Step) ToDomain() step.Step {
	return step.Step{
		ID:        s.ID,
		Payload:   s.Payload,
		Queue:     s.Queue,
		Status:    s.Status,
		StartedAt: s.StartedAt,
		StackID:   s.StackID,
		JobID:     s.JobID,
	}
}
