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

func (u *UseCases) CreateGroup(group entities.Group) error {
	return u.repo.CreateGroup(group)
}

func (u *UseCases) CreateAlbum(album entities.Album) error {
	return u.repo.CreateAlbum(album)
}

func (u *UseCases) CreateSong(song entities.Song) error {
	return u.repo.CreateSong(song)
}

func (u *UseCases) AllGroups() ([]entities.Group, error) {
	return u.repo.AllGroups()
}
