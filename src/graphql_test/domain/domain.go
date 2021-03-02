package domain

import "github.com/pajarraco93/graphql-test/src/graphql_test/domain/entities"

type Repository interface {
	Get(string) (entities.VideoGame, error)
	Create(entities.VideoGame) error
	ListAll() ([]entities.VideoGame, error)
}
