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

func (r *MySQLRepository) CreateAlbum(album entities.Album) error {
	query := fmt.Sprintf(`
		INSERT INTO 
			Albums (name, year, composedBy)
		VALUES 
			('%s', '%d', '%d')`,
		album.Name, album.Year, album.ComposedBy.ID,
	)
	_, err := r.engine.Query(query)
	return err
}

func (r *MySQLRepository) CreateSong(song entities.Song) error {
	query := fmt.Sprintf(`
		INSERT INTO 
			Songs (name, year, appearsIn)
		VALUES 
			('%s', '%d', '%d')`,
		song.Name, song.AppearsIn.ID,
	)
	_, err := r.engine.Query(query)
	return err
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

func (r *MySQLRepository) AllAlbums() (albums []entities.Album, err error) {
	query := fmt.Sprintf(`SELECT * FROM Albums`)
	rows, err := r.engine.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var album entities.Album
		err = rows.Scan(&album.ID, &album.Name, &album.Year, &album.ComposedBy.ID)
		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (r *MySQLRepository) GetGroupByID(ID int) (group entities.Group, err error) {
	query := fmt.Sprintf(`SELECT * FROM Groups WHERE groupID = '%d'`, ID)
	row := r.engine.QueryRow(query)

	err = row.Scan(&group.ID, &group.Name, &group.Genre)
	if err != nil {
		return group, err
	}

	return group, nil
}

func (r *MySQLRepository) GetAlbumByID(ID int) (album entities.Album, err error) {
	query := fmt.Sprintf(`SELECT * FROM Album WHERE id = '%d'`, ID)
	row := r.engine.QueryRow(query)

	err = row.Scan(&album.ID, &album.Name, &album.Year, &album.ComposedBy.ID)
	if err != nil {
		return album, err
	}

	return album, nil
}
