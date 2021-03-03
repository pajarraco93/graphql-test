package entities

import "github.com/google/uuid"

type Group struct {
	ID    int
	Name  string
	Genre string
}

type Album struct {
	ID         uuid.UUID
	Name       string
	ComposedBy Group
	Year       int
}

type Song struct {
	ID        uuid.UUID
	Name      string
	AppearsIn Album
}

type GenreType int

const (
	GenreTypeUnknown GenreType = iota
	GenreTypeRock
	GenreTypePop
	GenreTypeRnB
	GenreTypeMetal
)
