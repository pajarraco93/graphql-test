package entities

type VideoGame struct {
	Name  string
	Genre GenreType
}

type GenreType int

const (
	GenreTypeUnknown GenreType = iota
	GenreTypeShooter
	GenreTypeMOBA
	GenreTypeMMORPG
	GenreTypeAutochess
)
