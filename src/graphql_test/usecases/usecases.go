package usecases

import (
	"github.com/pajarraco93/graphql-test/src/graphql_test/domain"
	"github.com/pajarraco93/graphql-test/src/graphql_test/domain/entities"
)

type UseCases struct {
	repo domain.Repository
}

type UseCasesInterface interface {
	Get(string) (entities.VideoGame, error)
	Create(entities.VideoGame) error
	ListAll() ([]entities.VideoGame, error)
}

func NewUseCases(repo domain.Repository) UseCasesInterface {
	return &UseCases{
		repo: repo,
	}
}

func (u *UseCases) Get(string) (entities.VideoGame, error) {
	return entities.VideoGame{}, nil
}

func (u *UseCases) Create(entities.VideoGame) error {
	return nil
}

func (u *UseCases) ListAll() ([]entities.VideoGame, error) {
	return []entities.VideoGame{}, nil
}
