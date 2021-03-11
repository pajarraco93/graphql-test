package mysql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
)

var mockGroup = &entities.Group{
	ID:    1,
	Name:  "Iron Maiden",
	Genre: "metal",
}

var mockAlbum = &entities.Album{
	ID:         1,
	Name:       "Fear Of The Dark",
	Year:       1991,
	ComposedBy: *mockGroup,
}

var mockSong = &entities.Song{
	ID:        1,
	Name:      "Fear Of The Dark",
	AppearsIn: *mockAlbum,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAllGroups(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get all the groups", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Groups`)
			rows := sqlmock.NewRows([]string{"groupID", "name", "genre"}).
				AddRow(mockGroup.ID, mockGroup.Name, mockGroup.Genre)

			mock.ExpectQuery(query).WillReturnRows(rows)

			Convey("Then", func() {
				groups, err := repo.AllGroups()
				Convey("Groups must be retrieved and no error", func() {
					assert.NotNil(t, groups)
					assert.NoError(t, err)
				})
			})
		})
	})
}

func TestAllAlbums(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get all the albums", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Albums`)
			rows := sqlmock.NewRows([]string{"albumID", "name", "year", "composedBy"}).
				AddRow(mockAlbum.ID, mockAlbum.Name, mockAlbum.Year, mockAlbum.ComposedBy.ID)
			mock.ExpectQuery(query).WillReturnRows(rows)

			Convey("Then", func() {
				albums, err := repo.AllAlbums()
				Convey("Albums must be retrieved and no error", func() {
					assert.NotNil(t, albums)
					assert.NoError(t, err)
				})
			})
		})
	})
}

func TestAllSongs(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get all the songs", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Songs`)
			rows := sqlmock.NewRows([]string{"songID", "name", "appearsIn"}).
				AddRow(mockSong.ID, mockSong.Name, mockSong.AppearsIn.ID)
			mock.ExpectQuery(query).WillReturnRows(rows)

			Convey("Then", func() {
				songs, err := repo.AllSongs()
				Convey("Songs must be retrieved and no error", func() {
					assert.NotNil(t, songs)
					assert.NoError(t, err)
				})
			})
		})
	})
}

func TestGetGroupsByIDs(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get groups filtered by its ids", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Groups WHERE groupID IN (?)`)
			rows := sqlmock.NewRows([]string{"groupID", "name", "genre"}).
				AddRow(mockGroup.ID, mockGroup.Name, mockGroup.Genre).
				AddRow(2, "Muse", "alernative rock")
			mock.ExpectQuery(query).WithArgs("1,2").WillReturnRows(rows)

			Convey("Then", func() {
				groups, err := repo.GetGroupsByIDs([]int{1, 2})
				Convey("Groups must be retrieved and no error", func() {
					assert.NotNil(t, groups)
					assert.NoError(t, err)
					assert.Len(t, groups, 2)
				})
			})
		})
	})
}

func TestGetAlbumsByIDs(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get albums filtered by its ids", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Albums WHERE albumID IN (?)`)
			rows := sqlmock.NewRows([]string{"albumID", "name", "year", "composedBy"}).
				AddRow(mockAlbum.ID, mockAlbum.Name, mockAlbum.Year, mockAlbum.ComposedBy.ID).
				AddRow(2, "Piece Of Mind", 19882, mockAlbum.ComposedBy.ID)
			mock.ExpectQuery(query).WithArgs("1,2").WillReturnRows(rows)

			Convey("Then", func() {
				albums, err := repo.GetAlbumsByIDs([]int{1, 2})
				Convey("Albums must be retrieved and no error", func() {
					assert.NotNil(t, albums)
					assert.NoError(t, err)
					assert.Len(t, albums, 2)
				})
			})
		})
	})
}

func TestGetAlbumsByGroupID(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get albums filtered by its groupID", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Albums WHERE composedBy = ?`)
			rows := sqlmock.NewRows([]string{"albumID", "name", "year", "composedBy"}).
				AddRow(mockAlbum.ID, mockAlbum.Name, mockAlbum.Year, mockAlbum.ComposedBy.ID)
			mock.ExpectQuery(query).WithArgs(mockAlbum.ComposedBy.ID).WillReturnRows(rows)

			Convey("Then", func() {
				albums, err := repo.GetAlbumsByGroupID(mockAlbum.ID)
				Convey("Albums must be retrieved and no error", func() {
					assert.NotNil(t, albums)
					assert.NoError(t, err)
				})
			})
		})
	})
}

func TestGetSongsByAlbumID(t *testing.T) {
	db, mock := NewMock()
	repo := &MySQLRepository{
		engine: db,
	}

	Convey("Given we want to get songs filtered by its albumID", t, func() {
		Convey("When query is executed successfully", func() {
			query := regexp.QuoteMeta(`SELECT * FROM Songs WHERE appearsIn = ?`)
			rows := sqlmock.NewRows([]string{"songID", "name", "appearsIn"}).
				AddRow(mockSong.ID, mockSong.Name, mockSong.AppearsIn.ID)
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnRows(rows)

			Convey("Then", func() {
				songs, err := repo.GetSongsByAlbumID(mockSong.AppearsIn.ID)
				Convey("Songs must be retrieved and no error", func() {
					assert.NotNil(t, songs)
					assert.NoError(t, err)
				})
			})
		})
	})
}
