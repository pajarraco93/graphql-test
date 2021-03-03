package entities

import "github.com/google/uuid"

type Group struct {
	ID    uuid.UUID
	Name  string
	Genre GenreType
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
