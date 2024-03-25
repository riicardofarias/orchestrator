package job

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"orchestrator/internal/domain/job"
)

type Repository interface {
	Create(ctx context.Context, job job.Job) (string, error)
	UpdateStatus(ctx context.Context, id, status string) error
}

type repository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db:  db,
		col: db.Collection("jobs"),
	}
}

func (r *repository) Create(ctx context.Context, job job.Job) (string, error) {
	model := NewFromDomain(job)

	id, err := r.col.InsertOne(ctx, model)
	if err != nil {
		return "", err
	}

	return id.InsertedID.(string), nil
}

func (r *repository) UpdateStatus(ctx context.Context, id, status string) error {
	updateObj := bson.D{{"$set", bson.D{{"status", status}}}}

	_, err := r.col.UpdateByID(ctx, id, updateObj)
	if err != nil {
		return err
	}

	return nil
}
