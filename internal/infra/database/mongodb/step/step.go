package step

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"orchestrator/internal/domain/step"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (step.Step, error)
	Create(ctx context.Context, step step.Step) (string, error)
	UpdateStatus(ctx context.Context, id, status string) error
}

type repository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db:  db,
		col: db.Collection("steps"),
	}
}

func (r *repository) UpdateStatus(ctx context.Context, id, status string) error {
	updateObj := bson.D{{"$set", bson.D{{"status", status}}}}

	if _, err := r.col.UpdateByID(ctx, id, updateObj); err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(ctx context.Context, id string) (step.Step, error) {
	var model Step

	if err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return step.Step{}, nil
		}

		return step.Step{}, err
	}

	return model.ToDomain(), nil
}

func (r *repository) Create(ctx context.Context, step step.Step) (string, error) {
	model := NewFromDomain(step)

	id, err := r.col.InsertOne(ctx, model)
	if err != nil {
		return "", err
	}

	return id.InsertedID.(string), nil
}
