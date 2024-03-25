package stack

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"orchestrator/internal/domain/stack"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (stack.Stack, error)
	Create(ctx context.Context, stack stack.Stack) (string, error)
}

type repository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db:  db,
		col: db.Collection("stacks"),
	}
}

func (s *repository) FindByID(ctx context.Context, id string) (stack.Stack, error) {
	var model Stack

	err := s.col.FindOne(ctx, bson.M{"_id": id}).Decode(&model)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return stack.Stack{}, nil
		}

		return stack.Stack{}, err
	}

	return model.ToDomain(), nil
}

func (s *repository) Create(ctx context.Context, stack stack.Stack) (string, error) {
	model := NewFromDomain(stack)

	id, err := s.col.InsertOne(ctx, model)
	if err != nil {
		return "", err
	}

	return id.InsertedID.(string), nil
}
