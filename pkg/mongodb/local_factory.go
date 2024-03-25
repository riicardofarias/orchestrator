package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocalClientFactory struct {
	port string
}

func NewLocalClient() ClientFactory {
	return &LocalClientFactory{}
}

func (l *LocalClientFactory) GetClient(config ServerConfig) (*mongo.Client, error) {
	connString := fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client()
	opts.ApplyURI(connString).SetServerAPIOptions(serverAPI)

	return mongo.Connect(config.Context, opts)
}
