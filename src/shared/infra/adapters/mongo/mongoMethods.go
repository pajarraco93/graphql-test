package mongo

import "github.com/pajarraco93/graphql-test/src/graphql_test/domain/entities"

func (m *MongoRepository) Get(gameID string) (entities.VideoGame, error) {
	return entities.VideoGame{}, nil
}

func (m *MongoRepository) Find(term string) ([]entities.VideoGame, error) {
	return []entities.VideoGame{}, nil
}

func (m *MongoRepository) Create(game entities.VideoGame) error {
	return nil
}

func (m *MongoRepository) ListAll() ([]entities.VideoGame, error) {
	return []entities.VideoGame{}, nil
}
