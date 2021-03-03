package mysql

import "github.com/pajarraco93/graphql-test/pkg/library/domain/entities"

func (r *MySQLRepository) CreateGroup(entities.Group) error {
	return nil
}

func (r *MySQLRepository) CreateAlbum(entities.Album) error {
	return nil
}

func (r *MySQLRepository) CreateSong(entities.Song) error {
	return nil
}

func (r *MySQLRepository) AllGroups() ([]entities.Group, error) {
	return []entities.Group{}, nil
}
