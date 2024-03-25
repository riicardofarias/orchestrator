package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerClientFactory struct {
	port string
}

func NewServerClientFactory() ClientFactory {
	return &ServerClientFactory{}
}

func (l *ServerClientFactory) GetClient(config ServerConfig) (*mongo.Client, error) {
	connString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=%v&w=majority",
		config.User, config.Password, config.Host, config.RetryWrites,
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client()
	opts.ApplyURI(connString).SetServerAPIOptions(serverAPI)

	return mongo.Connect(config.Context, opts)
}
