package mysql

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"

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
		query := regexp.QuoteMeta(`SELECT * FROM Groups`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"groupID", "name", "genre"}).
				AddRow(mockGroup.ID, mockGroup.Name, mockGroup.Genre)

			mock.ExpectQuery(query).WillReturnRows(rows)

			Convey("Then", func() {
				groups, err := repo.AllGroups()
				Convey("Groups must be retrieved and no error", func() {
					So(groups, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				groups, err := repo.AllGroups()
				Convey("Error must be returned", func() {
					So(groups, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
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
		query := regexp.QuoteMeta(`SELECT * FROM Albums`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"albumID", "name", "year", "composedBy"}).
				AddRow(mockAlbum.ID, mockAlbum.Name, mockAlbum.Year, mockAlbum.ComposedBy.ID)
			mock.ExpectQuery(query).WillReturnRows(rows)

			Convey("Then", func() {
				albums, err := repo.AllAlbums()
				Convey("Albums must be retrieved and no error", func() {
					So(albums, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				albums, err := repo.AllAlbums()
				Convey("Error must be returned", func() {
					So(albums, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
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
		query := regexp.QuoteMeta(`SELECT * FROM Songs`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"songID", "name", "appearsIn"}).
				AddRow(mockSong.ID, mockSong.Name, mockSong.AppearsIn.ID)
			mock.ExpectQuery(query).WillReturnRows(rows)

			Convey("Then", func() {
				songs, err := repo.AllSongs()
				Convey("Songs must be retrieved and no error", func() {
					So(songs, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				songs, err := repo.AllSongs()
				Convey("Error must be returned", func() {
					So(songs, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
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
		query := regexp.QuoteMeta(`SELECT * FROM Groups WHERE groupID IN (?)`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"groupID", "name", "genre"}).
				AddRow(mockGroup.ID, mockGroup.Name, mockGroup.Genre).
				AddRow(2, "Muse", "alernative rock")
			mock.ExpectQuery(query).WithArgs("1,2").WillReturnRows(rows)

			Convey("Then", func() {
				groups, err := repo.GetGroupsByIDs([]int{1, 2})
				Convey("Groups must be retrieved and no error", func() {
					So(groups, ShouldNotBeNil)
					So(err, ShouldBeNil)
					So(groups, ShouldHaveLength, 2)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				groups, err := repo.GetGroupsByIDs([]int{1, 2})
				Convey("Error must be returned", func() {
					So(groups, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
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
		query := regexp.QuoteMeta(`SELECT * FROM Albums WHERE albumID IN (?)`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"albumID", "name", "year", "composedBy"}).
				AddRow(mockAlbum.ID, mockAlbum.Name, mockAlbum.Year, mockAlbum.ComposedBy.ID).
				AddRow(2, "Piece Of Mind", 19882, mockAlbum.ComposedBy.ID)
			mock.ExpectQuery(query).WithArgs("1,2").WillReturnRows(rows)

			Convey("Then", func() {
				albums, err := repo.GetAlbumsByIDs([]int{1, 2})
				Convey("Albums must be retrieved and no error", func() {
					So(albums, ShouldNotBeNil)
					So(err, ShouldBeNil)
					So(albums, ShouldHaveLength, 2)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				albums, err := repo.GetAlbumsByIDs([]int{1, 2})
				Convey("Error must be returned", func() {
					So(albums, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
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
		query := regexp.QuoteMeta(`SELECT * FROM Albums WHERE composedBy = ?`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"albumID", "name", "year", "composedBy"}).
				AddRow(mockAlbum.ID, mockAlbum.Name, mockAlbum.Year, mockAlbum.ComposedBy.ID)
			mock.ExpectQuery(query).WithArgs(mockAlbum.ComposedBy.ID).WillReturnRows(rows)

			Convey("Then", func() {
				albums, err := repo.GetAlbumsByGroupID(mockAlbum.ID)
				Convey("Albums must be retrieved and no error", func() {
					So(albums, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				albums, err := repo.GetAlbumsByGroupID(mockAlbum.ID)
				Convey("Error must be returned", func() {
					So(albums, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
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
		query := regexp.QuoteMeta(`SELECT * FROM Songs WHERE appearsIn = ?`)

		Convey("When query is executed successfully", func() {
			rows := sqlmock.NewRows([]string{"songID", "name", "appearsIn"}).
				AddRow(mockSong.ID, mockSong.Name, mockSong.AppearsIn.ID)
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnRows(rows)

			Convey("Then", func() {
				songs, err := repo.GetSongsByAlbumID(mockSong.AppearsIn.ID)
				Convey("Songs must be retrieved and no error", func() {
					So(songs, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When query fails", func() {
			mock.ExpectQuery(query).WithArgs(mockSong.AppearsIn.ID).WillReturnError(driver.ErrBadConn)

			Convey("Then", func() {
				songs, err := repo.GetSongsByAlbumID(mockSong.AppearsIn.ID)
				Convey("Error must be returned", func() {
					So(songs, ShouldBeEmpty)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
