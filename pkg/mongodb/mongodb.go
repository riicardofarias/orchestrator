package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	config  ServerConfig
	factory ClientFactory
}

type ClientFactory interface {
	GetClient(config ServerConfig) (*mongo.Client, error)
}

type Client interface {
	Database() *mongo.Database
	Close(ctx context.Context) error
	UseSession(ctx context.Context, fn func(sessionContext mongo.SessionContext) error) error
}

type ServerConfig struct {
	Context     context.Context
	Host        string
	Port        string
	User        string
	Password    string
	Database    string
	RetryWrites bool
}

type mongoClient struct {
	database *mongo.Database
	client   *mongo.Client
}

func NewMongoDB(config ServerConfig, factory ClientFactory) *MongoDB {
	return &MongoDB{
		config:  config,
		factory: factory,
	}
}

func (m *MongoDB) GetClient() (Client, error) {
	return newClient(m.config, m.factory)
}

func newClient(config ServerConfig, factory ClientFactory) (Client, error) {
	client, err := factory.GetClient(config)
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Database)

	return &mongoClient{
		database: db,
		client:   client,
	}, nil
}

func (m *mongoClient) Database() *mongo.Database {
	return m.database
}

func (m *mongoClient) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *mongoClient) UseSession(ctx context.Context, fn func(sessionContext mongo.SessionContext) error) error {
	if err := m.database.Client().UseSession(ctx, fn); err != nil {
		return err
	}

	return nil
}
