package job

import (
	"orchestrator/internal/domain/job"
	"time"
)

type Job struct {
	ID         string    `bson:"_id,omitempty"`
	StartedAt  time.Time `bson:"started_at"`
	FinishedAt time.Time `bson:"finished_at,omitempty"`
	Status     string    `bson:"status"`
	StackID    string    `bson:"stack_id"`
}

func NewFromDomain(job job.Job) *Job {
	return &Job{
		ID:         job.ID,
		StartedAt:  job.StartedAt,
		FinishedAt: job.FinishedAt,
		Status:     job.Status,
		StackID:    job.StackID,
	}
}
