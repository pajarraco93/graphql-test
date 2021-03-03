package mysql

import (
	"fmt"

	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
)

func (r *MySQLRepository) CreateGroup(group entities.Group) error {
	query := fmt.Sprintf(`
		INSERT INTO 
			Groups (name, genre)
		VALUES 
			('%s', '%s')`,
		group.Name, group.Genre,
	)
	_, err := r.engine.Query(query)
	return err
}

func (r *MySQLRepository) CreateAlbum(entities.Album) error {
	return nil
}

func (r *MySQLRepository) CreateSong(entities.Song) error {
	return nil
}

func (r *MySQLRepository) AllGroups() (groups []entities.Group, err error) {
	query := fmt.Sprintf(`SELECT * FROM Groups`)
	rows, err := r.engine.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var group entities.Group
		err = rows.Scan(&group.ID, &group.Name, &group.Genre)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}
