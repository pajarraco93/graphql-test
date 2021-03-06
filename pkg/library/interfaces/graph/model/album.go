package model

type Album struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ComposedBy int    `json:"composedBy"`
	Year       *int   `json:"year"`
}
