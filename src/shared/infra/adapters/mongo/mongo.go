package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pajarraco93/graphql-test/src/graphql_test/domain"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository() (domain.Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	repo := &MongoRepository{client}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = repo.connect(ctx)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (m *MongoRepository) connect(ctx context.Context) error {
	err := m.client.Connect(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoRepository) Disconnect() error {
	return m.client.Disconnect(context.Background())
}
