package usecases

import (
	"github.com/pajarraco93/graphql-test/pkg/library/domain"
	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
)

type UseCases struct {
	repo domain.Repository
}

type UseCasesInterface interface {
	CreateGroup(entities.Group) error
	CreateAlbum(entities.Album) error
	CreateSong(entities.Song) error
	AllGroups() ([]entities.Group, error)
}

func NewUseCases(repo domain.Repository) UseCasesInterface {
	return &UseCases{
		repo: repo,
	}
}

func (u *UseCases) CreateGroup(entities.Group) error {
	return nil
}

func (u *UseCases) CreateAlbum(entities.Album) error {
	return nil
}

func (u *UseCases) CreateSong(entities.Song) error {
	return nil
}

func (u *UseCases) AllGroups() ([]entities.Group, error) {
	return []entities.Group{}, nil
}
