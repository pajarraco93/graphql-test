package entities

type Group struct {
	ID    int
	Name  string
	Genre string
}

type Album struct {
	ID         int
	Name       string
	ComposedBy Group
	Year       int
}

type Song struct {
	ID        int
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
