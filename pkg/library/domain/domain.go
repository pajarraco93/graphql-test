package domain

import "github.com/pajarraco93/graphql-test/pkg/library/domain/entities"

type Repository interface {
	CreateGroup(entities.Group) error
	CreateAlbum(entities.Album) error
	CreateSong(entities.Song) error

	AllGroups() ([]entities.Group, error)
	GetGroupByName(string) (entities.Group, error)
}
